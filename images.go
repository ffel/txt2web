package txt2web

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ffel/pandocfilter"
)

const ImagePath = "images/"

// ImageNode analyse chunks for references to local images.  These
// will be copied into the target directory and the link will be
// updated accordingly

// type FuncProcessImage func(wg *sync.WaitGroup)

func Images(in <-chan Chunk) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		chunknr := 0

		var wg sync.WaitGroup

		for c := range in {
			chunknr++
			wg.Add(1)
			li := &localImages{chunknr, 0, make(map[string]string)}

			pandocfilter.Walk(li, c.Json)

			go func(l *localImages) {
				defer wg.Done()

				for target, orig := range l.renames {
					fmt.Printf("- copy %q to %q\n", orig, target)
				}
			}(li)

			out <- c
		}

		wg.Wait()
		close(out)
	}()

	return out
}

type localImages struct {
	chunknr int               // each chunk gets its own number to prevent clashes
	imgnr   int               // each image in a chunk gets its own number
	renames map[string]string // target - orig file name map
}

func (img *localImages) rename(path string) (string, bool) {
	u, err := url.Parse(path)

	if err != nil {
		log.Println("image path parse problem", err)
		return path, false
	}

	if u.Host == "" {
		img.imgnr++

		_, file := filepath.Split(u.Path)
		ext := filepath.Ext(file)
		pre := strings.TrimSuffix(file, ext)
		u.Path = fmt.Sprintf("%s%s_%d_%d%s", ImagePath, pre, img.chunknr, img.imgnr, ext)

		return u.String(), true
	}

	return path, false
}

func (img *localImages) Value(level int, key string, value interface{}) (bool, interface{}) {

	ok, t, c := pandocfilter.IsTypeContents(value)

	if ok && t == "Image" {
		path, err := pandocfilter.GetString(c, "1", "0")

		if err != nil {
			log.Println("image error", err)
		}

		targetname, local := img.rename(path)

		if err := pandocfilter.SetString(c, targetname, "1", "0"); err != nil {
			log.Println("image rename error", err)
		}

		if local {
			// targetnames are unique, and we have to assure that
			// every target will exist, so it is more safe to use
			// target as the key
			img.renames[targetname] = path
		}
	}

	return true, value
}
