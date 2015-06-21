package txt2web

import (
	"log"

	"github.com/ffel/pandocfilter"
)

// The references is the node in the pipeline that replaces markdown references
// (which work within markdown files) for references that work between html
// pages.  This is esp. usefull if the node contains contents that may be
// split up.
//
// This node is a "two subprocess" node.  The first sub-process finds the
// id's, the second sub-process translates references to those id's
//
// We assume you write ordinary pandoc txt documents with ordinary internal
// references.  There is only one exception: as soon as you want to refer to
// a section in a different txt file, you'll have to use the eventual reference.
//
// I'd like to stress, this only affects the `[reference](#link)`, not the
// id that can be given to a section (this is the ordinary default pandoc
// id or an overriden id to be used for internal references).

// RefChunk wraps Chunk with a ref translation map.  It is the intermeditate
// between the two subl processes.
type RefChunk struct {
	Chunk                          // original chunk
	Translations map[string]string // internal ref -> external ref
}

// References wraps two sub processes that together translate references.
func References(in <-chan Chunk) <-chan Chunk {
	return ref_translator(ref_finder(in))
}

func ref_finder(in <-chan Chunk) <-chan RefChunk {
	out := make(chan RefChunk)
	go func() {
		for c := range in {
			finder := &reffinder{make(map[string]string)}

			pandocfilter.Walk(finder, c.Json)

			// fmt.Printf("%#v\n", finder.refs)

			out <- RefChunk{c, finder.refs}
		}
		close(out)
	}()
	return out
}

func ref_translator(in <-chan RefChunk) <-chan Chunk {
	out := make(chan Chunk)
	go func() {
		for rc := range in {

			out <- rc.Chunk
		}
		close(out)
	}()
	return out
}

/*
Sub Process 1: sections have reference id's

    + "c": list:
        + "0": value: 2
        + "1": list:
            + "0": value: level-two
            + "1": list:
            + "2": list:
        + "2": list:
            + "0" - Str: "Level"
            + "1" - Space
            + "2" - Str: "Two"
    + "t": value: Header

Headers have `t` equal to "Header" and a list of three items as contents,
level, reference, and title.

For Subprocess 2

references have the following structure:

	+ "c": list:
	    + "0": list:
	        + "0" - Str: "Ref"
	        + "1" - Space
	        + "2" - Str: "text"
	    + "1": list:
	        + "0": value: #reference
	        + "1": value:
	+ "t": value: Link

That is, `t` is "Link" and `c` is a list with two values

*/

type reffinder struct {
	refs map[string]string
}

func (rf *reffinder) Value(level int, key string, value interface{}) (bool, interface{}) {

	ok, t, c := pandocfilter.IsTypeContents(value)

	if ok && t == "Header" {
		secLevel, err := pandocfilter.GetNumber(c, "0")

		if err != nil {
			log.Println(err)
			return false, value
		}

		if secLevel != 1 {
			return false, value
		}

		ref, err := pandocfilter.GetString(c, "1", "0")

		if err != nil {
			log.Println(err)
			return false, value
		}

		if _, exists := rf.refs[ref]; exists {
			log.Printf("reffinder - duplicate key: %v\n", ref)
		}

		rf.refs[ref] = "foobar"

		return false, value
	}

	return true, value
}
