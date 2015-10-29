package txt2web

import (
	"encoding/json"
	"log"

	"github.com/ffel/pandocfilter"
)

/*
OK, al die tijd zat ik er enorm dicht bij!!

Ik heb de foutmelding "2015/10/29 08:32:52 markdownChan:<nil>"
verkeerd geïnterpreteerd (bye bye lieve uren).

Het is veroorzaakt door pandoc die over de json struikelt.

De syntax van json is dus al die tijd fout geweest.

Het hele clone verhaal is dus niet nodig geweest.  Maar nu
ik ben gaan clonen, is het wel noodzakelijk om iets met de
return waarde van `pandocfilter.Walk(ul, c.Json)` te doen.

-----

Ik ben bezig om links aan de secties toe te voegen.

Ik heb code die de interne structuur van deze sectie maakt
(`wrapHeader`), hoewel ik niet zeker weet of de formaat van de header
precies klopt. Dit is wel te repareren.

Ik zie dat niets van de nieuwe sectie header in de uiteindelijke uitvoer
terecht komt (ik heb het over de tests in uplink\_test).

Het lijkt er op dat de nieuwe struct niet ingevoegd wordt.

Ik zie dat de return value in `pandocfilter.Walk()` niet wordt gebruikt.
Dat zou kunnen verklaren dat ik de oorspronkelijke data structuur nog
heb.

> Relevant is dat de meeste wijzigingen aan de structuur in txt2web via
> de `SetString()` en aanverwante functies is. Dat zijn (blijkbaar)
> in-place wijzigingen. In het geval van Uplink heb ik een ander soort
> wijziging.

Gisteren had ik een iets andere aanpak waarbij ik in interne structuur
aanpassingen deed.  Dit leverde "2015/10/29 08:32:52 markdownChan:<nil>"
fouten op (datum is onjuist).

Ik ben er van uit gegaan dat ik clonen van objecten moet gebruiken.

Ik heb nu een variant die een nieuw dataobject teruggeeft, maar op het
moment dat ik `c.Json = pandocfilter.Walk(ul, c.Json)` gebruik loop ik
weer op dezelfde "2015/10/29 08:32:52 markdownChan:<nil>" problemen.

*/

// UpLinkNode takes section headers and translates these into links
// to the containing sections.
func UpLinkNode(in <-chan Chunk) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		for c := range in {

			ul := &uplinks{}
			// Walk returns an interface{}, can we completely ignore that?
			// dit levert weer een foutmelding "2015/10/29 08:32:52 markdownChan:<nil>"
			// op ("dit" is de toevoeging van "c.Json =")

			// vooruit clonen heeft geen zin
			// clone, err := clone(c.Json)

			// if err == nil {
			// 	c.Json = pandocfilter.Walk(ul, clone)
			// } else {
			// 	log.Println("UpLinkNode", err)
			// }

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

			// hier lijkt het mis te gaan.  Het is wellicht de eerste keer dat ik met een ander object aan kom
			// tot nu toe vooral in-place wijzigingen (als al wijzigingen)...
			return false, wrapHeader(c, link)
		}

		return true, value
	}

	return true, value
}

func wrapHeader(header interface{}, link string) interface{} {
	// without cloning, setObject performs an inplace replacement which disturbs the channel
	orig, err := clone(header)

	// fmt.Printf("orig: %# v\n", pretty.Formatter(header))

	if err != nil {
		log.Printf("wrapHeader: %v\n", err)
		return header
	}

	// title, err := pandocfilter.GetObject(orig, "1")
	title, err := pandocfilter.GetSlice(orig, "2")

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

	err = pandocfilter.SetObject(orig, linkedHeader, "2")

	if err != nil {
		log.Printf("wrapHeader %v\n", err)
	}

	// controleer of je de juiste structuur terug krijgt ...

	// fmt.Printf("reworked header: %# v\n", pretty.Formatter(orig))

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
