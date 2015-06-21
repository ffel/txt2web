package txt2web

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ffel/piperunner"
)

const hello = `# hello

world!
`

func ExampleReferences() {
	// chunks := contentGen(hello)

	// for c := range chunks {
	// 	fmt.Printf("%#v\n", c.Json)
	// }

	markdownTerm(contentGen(hello))

	// output:
	// boo
}

// we need a util_test.go file with all test utilities.
// what I need is a generator that creates chunks by string
// contents

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
				log.Fatal(err)
			}

			var jsondata interface{}
			err := json.Unmarshal(result.Text, &jsondata)

			if err != nil {
				log.Fatal(err)
			}

			out <- Chunk{Json: jsondata, Section: i}
		}
		close(out)
	}()

	return out
}

// markdownTerm does basically the same as WriteHtml, except that is does
// produce markdown
func markdownTerm(in <-chan Chunk) {

	// je moet kiezen, of een terminator die print, of een node die een closure heeft

	for c := range in {

		bytes, err := json.Marshal(c.Json)

		if err != nil {
			log.Fatal(err)
		}

		resultc := piperunner.Exec("pandoc -f json -t markdown", bytes)

		result := <-resultc

		if result.Err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(result.Text))
		fmt.Println("---")
	}
}
