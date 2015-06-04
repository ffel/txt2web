package txt2web

import (
	"log"
	"path/filepath"
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
