package txt2web

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ffel/piperunner"
)

// namegen is a simple generator for Generator testing
func namegen(names ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, n := range names {
			out <- n
		}
		close(out)
	}()
	return out
}

// chunkterm is a simple chunk terminator
func chunkterm(in <-chan Chunk) {
	for c := range in {
		fmt.Println(c)
	}
}

// contentGen does basically the same as generator.go getJson,
// except that contentgen expects test string contents and terminates
// in case of an error.  (and contentGen acts as a pipeline generator)
func contentGen(content ...string) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		for i, c := range content {

			resultc := piperunner.Exec("pandoc -f markdown -t json", []byte(c))

			result := <-resultc

			if err := result.Err; err != nil {
				log.Fatal("contentGen:", err)
			}

			var jsondata interface{}
			err := json.Unmarshal(result.Text, &jsondata)

			if err != nil {
				log.Fatal("contentGen:", err)
			}

			out <- Chunk{Json: jsondata, Section: i}
		}
		close(out)
	}()

	return out
}

// setFile sets the file name in chunk
func setFile(in <-chan Chunk, filename string) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		for c := range in {
			c.Path = filename

			out <- c
		}
		close(out)
	}()

	return out
}

// markdownTerm does basically the same as WriteHtml, except that is does
// produce markdown
func markdownTerm(in <-chan Chunk) {
	for c := range in {

		bytes, err := json.Marshal(c.Json)

		if err != nil {
			log.Fatal("markdownTerm - marshal json:", err)
		}

		resultc := piperunner.Exec("pandoc -f json -t markdown", bytes)

		result := <-resultc

		if result.Err != nil {
			log.Println("markdownTerm: - result:", string(result.Text))
		}

		fmt.Print(string(result.Text))
		fmt.Println("---")
	}
}

// markdownChan proceduces markdown over a channel
func markdownChan(in <-chan Chunk) <-chan []byte {
	out := make(chan []byte)

	go func() {
		for c := range in {

			bytes, err := json.Marshal(c.Json)

			if err != nil {
				log.Fatal("markdownChan:", err)
			}

			resultc := piperunner.Exec("pandoc -f json -t markdown", bytes)

			result := <-resultc

			if result.Err != nil {
				log.Fatal("markdownChan:", err)
			}

			out <- result.Text
		}
		close(out)
	}()

	return out
}
