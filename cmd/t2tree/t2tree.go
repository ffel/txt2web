// t2tree is a dev tool that prints the pandoc parse tree from a txt file
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/ffel/pandocfilter"
	"github.com/ffel/txt2web"
)

// run - go run t2tree.go ../../example/*.txt

func main() {
	out := TreeNode(txt2web.Generate(filenames(os.Args[1:]...)))

	for tree := range out {
		fmt.Println(tree)
	}
}

// filenames creates a channel of strings out of names
func filenames(names ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, n := range names {
			out <- n
		}
		close(out)
	}()
	return out
}

// TreeNode takes chunks from the inbound channel and writes pandoc trees
func TreeNode(in <-chan txt2web.Chunk) <-chan string {
	out := make(chan string)

	go func() {
		for c := range in {
			tree := NewTree()
			pandocfilter.Walk(tree, c.Json)
			out <- tree.String()
		}
		close(out)
	}()

	return out
}

// copied from https://github.com/ffel/pandocfilter/blob/master/cmd/pdtree/tree.go

// NewTree initiates a Tree object and returns its pointer
func NewTree() *Tree {
	return &Tree{&bytes.Buffer{}}
}

// Tree prints a tree view of pandoc json, it also duplicates
// the pandoc json as does Duplicator
type Tree struct {
	buff *bytes.Buffer
}

func (t *Tree) Value(level int, key string, value interface{}) (bool, interface{}) {
	_, isList := value.([]interface{})

	if isList {
		fmt.Fprintf(t.buff, "%s+ %q: list:\n", t.indent(level), key)

		return true, nil
	}

	// a cstring is a special type of map
	isTC, tval, cval := pandocfilter.IsTypeContents(value)

	// don't do anything special with known collection types
	// as the returned values are used again by pandoc

	if isTC {
		switch tval {
		case pandocfilter.Space:
			fmt.Fprintf(t.buff, "%s+ %q - %s\n", t.indent(level), key, tval)
			return false, value
		case pandocfilter.Str:
			fmt.Fprintf(t.buff, "%s+ %q - %s: %q\n", t.indent(level), key, tval, cval.(string))
			return false, value
		}
	}

	_, isSet := value.(map[string]interface{})

	if isSet {
		fmt.Fprintf(t.buff, "%s+ %q: map:\n", t.indent(level), key)

		return true, nil
	}

	// value is not identifies as something special
	fmt.Fprintf(t.buff, "%s+ %q: value: %v\n", t.indent(level), key, value)

	return true, value
}

func (t *Tree) String() string {
	return t.buff.String()
}

func (t *Tree) indent(level int) string {
	return strings.Repeat("    ", level)
}
