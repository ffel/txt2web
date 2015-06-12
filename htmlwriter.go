package txt2web

// the htmlwriter is a pipeline (pseudo) terminator that translates the input
// to html files

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/ffel/piperunner"
)

// WriteHtml writes Chunks to HtmlFile
func WriteHtml(in <-chan Chunk) <-chan HtmlFile {

	htmlfilec := make(chan HtmlFile)

	go func() {

		// we need a waitgroup to assure that `htmlfilec` is closed
		// only after all inner go routines have been closed.

		var wg sync.WaitGroup

		for chunk := range in {
			wg.Add(1)
			// use a closure here to prevent blocking

			// this inner closure needs to receive chunk as an argument
			// to prevent race conditions due to the typical closure issue...
			go func(c Chunk) {
				defer wg.Done()
				bytes, err := json.Marshal(c.Json)

				if err != nil {
					log.Println(err)
					return
				}

				// as soon as we need wrapping:
				// wrapped := []byte("[ { \"unMeta\" : {  } },")
				// wrapped = append(wrapped, bytes...)
				// wrapped = append(wrapped, []byte("]")...)

				resultc := piperunner.Exec("pandoc -f json -t html", bytes)

				result := <-resultc

				if err = result.Err; err != nil {
					log.Println(err)
					return
				}

				htmlfilec <- HtmlFile{Path: c.Path, Contents: result.Text}

			}(chunk)
		}

		wg.Wait()
		close(htmlfilec)
	}()

	return htmlfilec
}
