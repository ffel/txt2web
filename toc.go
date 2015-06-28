package txt2web

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/ffel/pandocfilter"
)

// the toc is a node that creates a toc.  It can be used instead of
// the splitter

// the simplest is to print from within the walker, this makes it possible
// to use this link inside a chain.  The other alternative is to
// collect the data and either make Toc print the contents or send via
// a channel.

// Toc creates a table of contents
func Toc(in <-chan Chunk, writer io.Writer) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		for c := range in {

			t := &toc{}

			fmt.Printf("%v:\n", c.Path)

			pandocfilter.Walk(t, c.Json)

			for _, tocline := range t.sections {
				line := fmt.Sprintf("%s- %v", strings.Repeat("  ", tocline.level-1), tocline)

				if tocline.level == 1 {
					fmt.Fprintf(writer, "%80s\n%s\n", "#"+Webkey(c.Path, tocline.anchor), line)
				} else {
					fmt.Fprintln(writer, line)
				}
			}

			out <- c
		}
		close(out)
	}()

	return out
}

type tocEntry struct {
	level  int    // section level
	title  string // title
	anchor string // pandoc reference
	web    string // inter-page web reference
}

func (t tocEntry) String() string {
	return fmt.Sprintf("[%s](#%s)", t.title, t.anchor)
}

type toc struct {
	sections []tocEntry
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

			toc.sections = append(toc.sections, tocEntry{
				level:  int(secLevel),
				title:  strings.TrimSpace(col.value),
				anchor: id,
			})
		}

		return false, value
	}

	return true, value
}
