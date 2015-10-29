package txt2web

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/ffel/piperunner"
	"github.com/kr/pretty"
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
func setFiles(in <-chan Chunk, filenames ...string) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		i := 0
		for c := range in {
			if i < len(filenames) {
				c.Path = filenames[i]
			} else if len(filenames) > 0 {
				c.Path = filenames[len(filenames)-1]
			}
			i++

			out <- c
		}
		close(out)
	}()

	return out
}

// structPrint can be included in the chain to pretty print the currect internal data
func structPrint(in <-chan Chunk, prefix string, on bool) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		for c := range in {
			if on {
				fmt.Printf("\n%v\n%# v\n", prefix, pretty.Formatter(c.Json))
			}
			out <- c
		}
		close(out)
	}()

	return out
}

// jsonPrint pretty prints the internal data structure
func jsonPrint(jsonstring string) string {

	var jsondata interface{}
	err := json.Unmarshal([]byte(jsonstring), &jsondata)

	if err != nil {
		log.Fatal("jsonPrint:", err)
	}

	return fmt.Sprintf("%# v\n", pretty.Formatter(jsondata))
}

// markdownTerm terminates a pipeline and prints chunks as a sorted list of markdown
func markdownTerm(in <-chan Chunk, full bool) {
	var results []markdownFile

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

		results = append(results, markdownFile{path: c.Path, markdown: string(result.Text)})
	}

	sort.Sort(markdownByTitle(results))

	for i, f := range results {
		if full {
			fmt.Printf("<<%d - %q>>\n%v", i, f.path, f.markdown)
		} else {
			fmt.Printf("%d. %q\n", i, f.path)
		}
	}
}

type markdownFile struct {
	path     string
	markdown string
}

type markdownByTitle []markdownFile

func (m markdownByTitle) Len() int {
	return len(m)
}

func (m markdownByTitle) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m markdownByTitle) Less(i, j int) bool {
	return m[i].path < m[j].path
}

// markdownChan proceduces markdown over a channel
func markdownChan(in <-chan Chunk) <-chan []byte {
	out := make(chan []byte)

	go func() {
		for c := range in {

			bytes, err := json.Marshal(c.Json)

			if err != nil {
				log.Fatal("markdownChan: ", err)
			}

			resultc := piperunner.Exec("pandoc -f json -t markdown", bytes)

			result := <-resultc

			if result.Err != nil {
				log.Fatalf("markdownChan: %v\n%# v\n-- pandoc could not convert this json\n", err, pretty.Formatter(c.Json))
			}

			out <- result.Text
		}
		close(out)
	}()

	return out
}
