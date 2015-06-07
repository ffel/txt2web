package txt2web

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ffel/piperunner"
)

// the htmlwriter is a pipeline terminator that translates the input
// to html files

func WriteHtml(in <-chan Chunk) {
	for chunk := range in {
		// use a closure here to prevent blocking
		go func() {
			bytes, err := json.Marshal(chunk.Json)

			if err != nil {
				log.Println(err)
				return
			}

			// as soon as we need wrapping:
			// wrapped := []byte("[ { \"unMeta\" : {  } },")
			// wrapped = append(wrapped, bytes...)
			// wrapped = append(wrapped, []byte("]")...)

			resultc := piperunner.Exec("pandoc -f json -t markdown", bytes)

			result := <-resultc

			if err = result.Err; err != nil {
				log.Println(err)
				return
			}

			fmt.Printf("%v:\n%v\n\n", chunk.Title, string(result.Text))

		}()
	}
}
