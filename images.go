package txt2web

import (
	"fmt"
	"log"

	"github.com/ffel/pandocfilter"
)

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
		for c := range in {

			li := &localImages{}

			pandocfilter.Walk(li, c.Json)

			out <- c
		}
		close(out)
	}()

	return out
}

/*
+ "0": map:
    + "c": list:
        + "0": list:
            + "0" - Str: "image"
            + "1" - Space
            + "2" - Str: "description"
        + "1": list:
            + "0": value: image.png
            + "1": value: fig:
    + "t": value: Image
*/

type localImages struct {
}

func (img *localImages) Value(level int, key string, value interface{}) (bool, interface{}) {

	ok, t, c := pandocfilter.IsTypeContents(value)

	if ok && t == "Image" {
		_ = c
		path, err := pandocfilter.GetString(c, "1", "0")

		if err != nil {
			log.Println("image error", err)
		} else {
			fmt.Printf("found image link: %q\n", path)
		}
	}

	return true, value
}
