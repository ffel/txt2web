package txt2web

import (
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

			// determine parent - this is info needed to change c
			if len(ul.tocstack) == 0 {
				//fmt.Printf("section %q (%v) will refer to index.html\n", id, secLevel)
			} else {
				// fmt.Printf("section %q (%v) will refer to %q\n", id, secLevel,
				// 	ul.tocstack[len(ul.tocstack)-1])
			}

			// fmt.Printf("%# v\n", pretty.Formatter(c))

			// push current section (add "" for absent sections)
			if len(ul.tocstack) < secLevel {
				ul.tocstack = append(ul.tocstack, make([]string, secLevel-len(ul.tocstack))...)
			}
			ul.tocstack[len(ul.tocstack)-1] = id
		}

		return false, wrapHeader()
	}

	return true, value
}

func wrapHeader() interface{} {
	return []interface{}{
		float64(1),
		[]interface{}{
			"h2",
			[]interface{}{},
			[]interface{}{},
		},
		[]interface{}{
			map[string]interface{}{
				"t": "Link",
				"c": []interface{}{
					[]interface{}{
						map[string]interface{}{
							"t": "Str",
							"c": "h2",
						},
					},
					[]interface{}{
						"#h1",
						"",
					},
				},
			},
		},
	}
}

/*
func WrapTMath(typeMath, math string) interface{} {
	// explicit struct is possible, even simpler if we use
	// jmap and jslice aliases for map[string]interface{} and
	// []interface{}
	m := map[string]interface{}{
		"c": []interface{}{
			map[string]interface{}{
				"c": []interface{}{},
				"t": typeMath,
			},
			math,
		},
		"t": "Math",
	}

	return m
}
*/
