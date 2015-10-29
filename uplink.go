package txt2web

import (
	"encoding/json"
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

			c.Json = pandocfilter.Walk(ul, c.Json)

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

			// get rid of irrelevant part of the stack
			if len(ul.tocstack) >= secLevel {
				ul.tocstack = ul.tocstack[:secLevel-1]
			}

			// get rid of trailing "" elements
			for {
				last := len(ul.tocstack) - 1
				if last >= 0 && ul.tocstack[last] == "" {
					ul.tocstack = ul.tocstack[:last]
				} else {
					break
				}
			}

			link := ""

			// determine parent - this is info needed to change c
			if len(ul.tocstack) == 0 {
				link = "index.html"
			} else {
				link = ul.tocstack[len(ul.tocstack)-1]
			}

			// this is a major help in analysing the internal data structure
			// fmt.Printf("%# v\n", pretty.Formatter(c))

			// push current section (add "" for absent sections)
			if len(ul.tocstack) < secLevel {
				ul.tocstack = append(ul.tocstack, make([]string, secLevel-len(ul.tocstack))...)
			}
			ul.tocstack[len(ul.tocstack)-1] = id

			return false, wrapHeader(value, link)
		}

		return true, value
	}

	return true, value
}

func wrapHeader(header interface{}, link string) interface{} {

	// fmt.Printf("header: %# v\n", pretty.Formatter(header))

	title, err := pandocfilter.GetSlice(header, "c", "2")

	if err != nil {
		log.Printf("wrapHeader %v\n", err)
		return header
	}

	// fmt.Printf("title: %# v\n", pretty.Formatter(title))

	linkedHeader := []interface{}{
		map[string]interface{}{
			"t": "Link",
			"c": []interface{}{
				// []interface{}{
				title, // title inserted here
				// },
				[]interface{}{
					link, // link inserted here
					"",
				},
			},
		},
	}

	// fmt.Printf("linked header: %# v\n", pretty.Formatter(linkedHeader))

	err = pandocfilter.SetObject(header, linkedHeader, "c", "2")

	if err != nil {
		log.Printf("wrapHeader %v\n", err)
	}

	// fmt.Printf("reworked header: %# v\n", pretty.Formatter(header))

	return header
}

// make a new clone, lazy man implementation: convert to and from json
func clone(in interface{}) (interface{}, error) {

	bytes, err := json.Marshal(in)

	if err != nil {
		return in, err
	}

	var out interface{}
	err = json.Unmarshal(bytes, &out)

	if err != nil {
		return in, err
	}

	return out, nil
}
