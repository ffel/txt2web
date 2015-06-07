package txt2web

// the generator is the node in the pipeline that will create
// the very first Chunks.  Each chunk contains an entire file
// at this point

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ffel/piperunner"
)

// Generate is the chunk generator, one Chunk will contain one
// entire file at this point.
func Generate(filenames <-chan string) <-chan Chunk {
	chunks := make(chan Chunk)

	go func() {
		for f := range filenames {

			data, err := getJson(filepath.Join(txtRoot, f))

			if err != nil {
				log.Println(err)
				continue
			}

			// chunk is not yet a full fledged chunk
			chunks <- Chunk{Json: data, Path: f, Title: f, PandocId: "filechunk"}
		}
		close(chunks)
	}()

	return chunks
}

// function getJson converts contents of file to pandoc json
func getJson(file string) (interface{}, error) {
	txt, err := ioutil.ReadFile(file)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	resultc := piperunner.Exec("pandoc -f markdown -t json", txt)

	result := <-resultc

	if result.Err != nil {
		log.Println(err)
		return nil, result.Err
	}

	var jsondata interface{}
	err = json.Unmarshal(result.Text, &jsondata)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return jsondata, nil
}
