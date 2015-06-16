package txt2web

// the splitter is the node that is splits initial chunk that contain
// the text in one file into separate one-section-per-chunk chunks.
// each chunk gets the meta header of the original file.

import (
	"fmt"
	"log"
	"strings"

	"github.com/ffel/pandocfilter"
)

// Split splits chunks received on in in chunks with one # sections
func Split(in <-chan Chunk) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		for c := range in {
			spl := &splitter{}

			pandocfilter.Walk(spl, c.Json)

			for i, s := range spl.sections {
				out <- Chunk{
					Json:     wrapSection(spl.meta, s.contents),
					Path:     c.Path,
					Section:  i + 1,
					Title:    s.title,
					PandocId: s.id,
				}
			}

		}
		close(out)
	}()

	return out
}

// wrapSection makes piece of json valid pandoc json again
// it essentially adds the file yaml header in every section
func wrapSection(meta, content interface{}) interface{} {

	// use an anonymous struct
	// see, e.g. https://talks.golang.org/2012/10things.slide#2

	// btw, in earliers solutions, the wrapper was added to the json string
	// which makes it easier to get the `unMeta` correct (instead of `UnMeta`)

	return []interface{}{
		meta,
		content,
	}
}

/*
According to https://github.com/ffel/pandocfilter/blob/master/modify_test.go,
the following pandoc contents

	Hello
	=====

	world!

has the following json tree:

	[ { "unMeta" : {  } },
	  [ { "c" : [ 1,
	          [ "hello",
	            [  ],
	            [  ]
	          ],
	          [ { "c" : "Hello",
	              "t" : "Str"
	            } ]
	        ],
	      "t" : "Header"
	    },
	    { "c" : [ { "c" : "world!",
	            "t" : "Str"
	          } ],
	      "t" : "Para"
	    }
	  ]
	]

That is, a two element array with meta data and contents.  The contents
on its turn is again an array of elements.

Typical headers come in at level 2.  (I can only think of headers in block
quotes that come in on a higher level.  We can ignore these for now.)
*/

type sectiondata struct {
	title    string
	id       string
	contents []interface{}
}

func (s sectiondata) String() string {
	return fmt.Sprintf("[%s](%s)", s.title, s.id)
}

// splitter splits one chunk into sections
type splitter struct {
	sections []sectiondata // the collection of sections in a file
	meta     interface{}   // file meta data
}

func (spl *splitter) Value(level int, key string, value interface{}) (bool, interface{}) {

	ismeta, meta := isMeta(value)

	if ismeta {
		spl.meta = meta
	}

	// we can prevent deeper traversal through the tree by returning false
	// in case the level is 2
	if level == 2 {

		ok, t, c := pandocfilter.IsTypeContents(value)

		// check whether we have to start with the next section
		if ok && t == "Header" {
			secLevel, err := pandocfilter.GetNumber(c, "0")

			if err != nil {
				log.Println("Splitter error:", err)
				return false, value
			}

			// only interested in first order # sections
			if secLevel == 1 {
				spl.nextSection(c)
			}
		}

		if ok {
			current := spl.sections[len(spl.sections)-1]
			current.contents = append(current.contents, value)
			spl.sections[len(spl.sections)-1] = current
		}

		return false, value
	}

	return true, value
}

// nextSection prepares the next section
func (spl *splitter) nextSection(header interface{}) {
	id, err := pandocfilter.GetString(header, "1", "0")

	if err != nil {
		log.Println("Splitter error, could not determine header ID:", err)
		return
	}

	title, err := pandocfilter.GetObject(header, "2")

	if err != nil {
		log.Println("Splitter error, could not determine title:", err)
		return
	}

	col := &collector{}

	pandocfilter.Walk(col, title)

	spl.sections = append(spl.sections,
		sectiondata{
			id:    id,
			title: strings.TrimSpace(col.value),
		})
}

// isMeta determines if value is the meta structure and returns the
// complete meta set (including the "unMeta" tag)
func isMeta(value interface{}) (bool, interface{}) {
	set, isSet := value.(map[string]interface{})
	if !isSet {
		return false, nil
	}
	if len(set) != 1 {
		return false, nil
	}

	_, ok := set["unMeta"]
	if !ok {
		return false, nil
	}

	return true, set
}

// collector walks the header c and collects the Str
type collector struct {
	value string
}

func (coll *collector) Value(level int, key string, value interface{}) (bool, interface{}) {
	ok, t, c := pandocfilter.IsTypeContents(value)

	if ok && t == "Str" {
		coll.value += c.(string) + " "
	}

	return true, value
}
