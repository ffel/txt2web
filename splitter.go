package txt2web

// the splitter is the node that is splits initial chunk that contain
// the text in one file into separate one-section-per-chunk chunks.
// each chunk gets the meta header of the original file.

import (
	"fmt"
	"log"
	"path/filepath"
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
					Path:     changePath(c.Path, s.id),
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

// changePath replaces the file based chunk path for a section based
// file path
func changePath(orig, id string) string {
	dir := filepath.Dir(orig)
	return filepath.Join(dir, id+".html")
}

// wrapSection makes piece of json valid pandoc json again
// it essentially adds the file yaml header in every section
func wrapSection(meta, content interface{}) interface{} {
	return []interface{}{
		meta,
		content,
	}
}

/*
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
			if len(spl.sections) > 0 {
				current := spl.sections[len(spl.sections)-1]
				current.contents = append(current.contents, value)
				spl.sections[len(spl.sections)-1] = current
			} else {
				log.Println("Ignore contents before the first section.")
			}
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
