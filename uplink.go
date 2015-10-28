package txt2web

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ffel/pandocfilter"
	"github.com/kr/pretty"
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

			link := ""

			// determine parent - this is info needed to change c
			if len(ul.tocstack) == 0 {
				link = "index.html"
				//fmt.Printf("section %q (%v) will refer to index.html\n", id, secLevel)
			} else {
				link = ul.tocstack[len(ul.tocstack)-1]
				// fmt.Printf("section %q (%v) will refer to %q\n", id, secLevel,
				// 	ul.tocstack[len(ul.tocstack)-1])
			}

			// this is a major help in analysing the internal data structure
			// fmt.Printf("%# v\n", pretty.Formatter(c))

			// push current section (add "" for absent sections)
			if len(ul.tocstack) < secLevel {
				ul.tocstack = append(ul.tocstack, make([]string, secLevel-len(ul.tocstack))...)
			}
			ul.tocstack[len(ul.tocstack)-1] = id

			return false, wrapHeader(c, link)
		}

		return true, value
	}

	return true, value
}

func wrapHeader(header interface{}, link string) interface{} {
	// without cloning, setObject performs an inplace replacement which disturbs the channel
	orig, err := clone(header)

	fmt.Printf("link: %q\n", link)

	if err != nil {
		log.Printf("wrapHeader: %v\n", err)
		return header
	}

	title, err := pandocfilter.GetObject(orig, "1")

	if err != nil {
		log.Printf("wrapHeader %v\n", err)
		return header
	}

	linkedHeader := []interface{}{
		map[string]interface{}{
			"t": "Link",
			"c": []interface{}{
				[]interface{}{
					title, // title inserted here
				},
				[]interface{}{
					link, // link inserted here
					"",
				},
			},
		},
	}

	// hier lijkt iets mis te gaan, kijk maar eens naar origineel en gefabriceerde hier onder
	err = pandocfilter.SetObject(orig, linkedHeader, "2")

	if err != nil {
		log.Printf("wrapHeader %v\n", err)
	}

	fmt.Printf("reworked header: %# v\n", pretty.Formatter(orig))

	return orig
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

/*

[]interface {}{
    float64(1),
    []interface {}{
        "h2",
        []interface {}{
        },
        []interface {}{
        },
    },
    []interface {}{
        map[string]interface {}{
            "t": "Link",
            "c": []interface {}{
                []interface {}{
                    map[string]interface {}{
                        "t": "Str",
                        "c": "h2",
                    },
                },
                []interface {}{
                    "#h1",
                    "",
                },
            },
        },
    },
}

[]interface {}{
    float64(1),
    []interface {}{
        "main-ii",
        []interface {}{
        },
        []interface {}{
        },
    },
    []interface {}{
        map[string]interface {}{
            "c": []interface {}{
                []interface {}{
                    []interface {}{
                        "main-ii",
                        []interface {}{
                        },
                        []interface {}{
                        },
                    },
                },
                []interface {}{
                    "index.html",
                    "",
                },
            },
            "t": "Link",
        },
    },
}


*/
