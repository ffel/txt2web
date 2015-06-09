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
	// we need a waitgroup to make sure that the innen goroutine
	// of the very last item that is received on "in" completes
	// before the "htmlfilec" channel is closed.
	// In case we do not, the last result is never send on
	// the htmlfilec channel because htmlfilec is closed too
	// early (namely immediately after the last inner goroutine has
	// started)
	//
	// however, see http://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/
	var wg sync.WaitGroup
	go func() {
		for chunk := range in {
			wg.Add(1)
			// use a closure here to prevent blocking
			go func() {
				defer wg.Done()
				bytes, err := json.Marshal(chunk.Json)

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

				htmlfilec <- HtmlFile{Path: chunk.Path, Contents: result.Text}

			}()
		}

		wg.Wait()
		close(htmlfilec)
	}()

	return htmlfilec
}
