package txt2web

import (
	"fmt"
	"log"

	"github.com/ffel/pandocfilter"
)

// the toc is a node that creates a toc.  It can be used instead of
// the splitter

// the simplest is to print from within the walker, this makes it possible
// to use this link inside a chain.  The other alternative is to
// collect the data and either make Toc print the contents or send via
// a channel.

// Toc creates a table of contents
func Toc(in <-chan Chunk) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		for c := range in {

			t := &toc{}

			pandocfilter.Walk(t, c.Json)

			out <- c
		}
		close(out)
	}()

	return out
}

// splitter splits one chunk into sections
type toc struct {
	// sections []sectiondata // the collection of sections in a file
	// meta     interface{}   // file meta data
}

func (toc *toc) Value(level int, key string, value interface{}) (bool, interface{}) {

	// we can prevent deeper traversal through the tree by returning false
	// in case the level is 2
	if level == 2 {

		ok, t, c := pandocfilter.IsTypeContents(value)

		if ok && t == "Header" {
			secLevel, err := pandocfilter.GetNumber(c, "0")

			if err != nil {
				log.Println("toc error:", err)
				return false, value
			}

			id, err := pandocfilter.GetString(c, "1", "0")

			if err != nil {
				log.Println("toc error, could not determine header ID:", err)
				return false, value
			}

			title, err := pandocfilter.GetObject(c, "2")

			if err != nil {
				log.Println("toc error, could not determine title:", err)
				return false, value
			}

			col := &collector{}

			pandocfilter.Walk(col, title)

			fmt.Println("toc", secLevel, id, col.value)
		}

		return false, value
	}

	return true, value
}
