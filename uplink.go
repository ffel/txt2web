package txt2web

import (
	"fmt"
	"log"

	"github.com/ffel/pandocfilter"
)

// UpLinkNode takes section headers and translates these into links
// to the containing sections.
func UpLinkNode(in <-chan Chunk) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		for c := range in {

			ul := &uplinks{}
			pandocfilter.Walk(ul, c.Json)

			out <- c
		}
		close(out)
	}()

	return out
}

type uplinks struct {
	// maintain a stack of links here ...
	tocstack []string
}

func (ul *uplinks) Value(level int, key string, value interface{}) (bool, interface{}) {

	// we can prevent some unneccesary deep traversal here ...
	if level == 2 {
		ok, t, c := pandocfilter.IsTypeContents(value)

		if ok && t == "Header" {

			secLevelFloat, err := pandocfilter.GetNumber(c, "0")

			if err != nil {
				log.Println("uplink section level error:", err)
				return false, value
			}

			secLevel := int(secLevelFloat)

			id, err := pandocfilter.GetString(c, "1", "0")

			if err != nil {
				log.Println("uplink section id error:", err)
				return false, value
			}

			if len(ul.tocstack) == 0 {
				fmt.Printf("section %q (%v) will refer to index.html\n", id, secLevel)
			} else if len(ul.tocstack) >= secLevel-1 {
				fmt.Printf("section %q (%v) will refer to %q\n", id, secLevel,
					ul.tocstack[secLevel-2])
			}

			// handles nicely ordered contents

			if secLevel == len(ul.tocstack)+1 {
				ul.tocstack = append(ul.tocstack, id)
			} else if secLevel == len(ul.tocstack) {
				fmt.Println("**replace last value", id)
				ul.tocstack[len(ul.tocstack)-1] = id
			} else if secLevel == len(ul.tocstack)-1 {
				ul.tocstack = ul.tocstack[:len(ul.tocstack)-1]
				ul.tocstack[len(ul.tocstack)-1] = id
			}

			fmt.Printf("\t%#v\n", ul.tocstack)
		}

		return false, value
	}

	return true, value

}
