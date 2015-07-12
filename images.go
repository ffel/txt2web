package txt2web

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/ffel/pandocfilter"
)

const ImagePath = "images/"

// ImageNode analyse chunks for references to local images.  These
// will be copied into the target directory and the link will be
// updated accordingly

// Als ik dit testbaar wil, dan zal ik op de een of andere manier er
// voor moeten zorgen dat het behandelen van een foto niet hard is.
//
// Een manier is een extra out-channel met bestandsnamen, een andere
// manier is een extra functie als argument die wordt geroepen met
// de informatie als argument
//
// Het is nu vooral iets met de bestaande en de nieuwe naam van het
// plaatje, maar het zou wel heel prettig zijn wanneer je zeker
// weet of het bestand bestaat. (maar is dat een voorwaarde, ook de
// verwerkende functie kan een foutmelding geven)
//
// Een bijkomend probleem is dat out pas sluit wanneer zeker is
// dat alle foto's klaar zijn.

func Images(in <-chan Chunk) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		chunknr := 0

		for c := range in {

			chunknr++

			li := &localImages{chunknr, 0}

			pandocfilter.Walk(li, c.Json)

			out <- c
		}
		close(out)
	}()

	return out
}

type localImages struct {
	chunknr int // each chunk gets its own number to prevent clashes
	imgnr   int // each image in a chunk gets its own number
}

func (img *localImages) rename(path string) string {
	u, err := url.Parse(path)

	if err != nil {
		log.Println("image path parse problem", err)
		return path
	}

	if u.Host == "" {
		img.imgnr++

		_, file := filepath.Split(u.Path)
		ext := filepath.Ext(file)
		pre := strings.TrimSuffix(file, ext)
		u.Path = fmt.Sprintf("%s%s_%d_%d%s", ImagePath, pre, img.chunknr, img.imgnr, ext)

		return u.String()
	}

	return path
}

func (img *localImages) Value(level int, key string, value interface{}) (bool, interface{}) {

	ok, t, c := pandocfilter.IsTypeContents(value)

	if ok && t == "Image" {
		path, err := pandocfilter.GetString(c, "1", "0")

		if err != nil {
			log.Println("image error", err)
		}

		if err := pandocfilter.SetString(c, img.rename(path), "1", "0"); err != nil {
			log.Println("image rename error", err)
		}
	}

	return true, value
}
