package txt2web

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ffel/pandocfilter"
)

// Index is the node that analyzes chunks for syblings and childs and will
// add index chunks to the stream

// indexInfo is a node in a tree
type indexInfo struct {
	dir      string       // dir name (one level of path)
	sections []tocEntry   // section in current dir
	subdirs  []*indexInfo // subdirectories, nodes in the tree
}

func (ii *indexInfo) knowsDir(dir string) (bool, *indexInfo) {
	for _, d := range ii.subdirs {
		if d.dir == dir {
			return true, d
		}
	}
	return false, &indexInfo{}
}

func (ii *indexInfo) String() string {
	result := fmt.Sprintf("%s (%d): [", ii.dir, len(ii.sections))

	for _, c := range ii.subdirs {
		result += c.String() + ", "
	}

	return result + " ]"
}

// Index analyzes the in stream and adds index chunks to the stream
func Index(in <-chan Chunk) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		root := &indexInfo{dir: "."}

		for c := range in {

			path := filepath.Dir(c.Path)

			current := root

			// if we're not in root, we have to determine current as a node
			// in the tree
			if path != "." {

				elem := strings.Split(path, "/")

				// walk the tree and add sub dirs not yet in the tree
				for _, dir := range elem {

					if ok, sub := current.knowsDir(dir); ok {
						current = sub
					} else {
						current.subdirs = append(current.subdirs, &indexInfo{dir: dir})
						current = current.subdirs[len(current.subdirs)-1]
					}
				}
			}

			// add section data to current
			t := &toc{}

			pandocfilter.Walk(t, c.Json)

			current.sections = append(current.sections, t.sections...)

			// fmt.Printf("tree: %v\n", root)

			out <- c
		}

		// we need a wait group because pandoc is invoked which makes
		// the result asynchronous
		var wg sync.WaitGroup
		addIndex(out, root, "", wg)
		wg.Wait()

		close(out)
	}()

	return out
}

func addIndex(out <-chan Chunk, node *indexInfo, path string, wg sync.WaitGroup) {

	path = filepath.Join(path, node.dir)

	go func() {
		// there is a slight chance that this goroutine terminates before
		// the recursive call is invoked - this could terminate the waitgroup
		// to soon. I'm not sure that depth first is a solution either ...
		wg.Add(1)

		fmt.Printf("dir %q has sections:\n", path)

		for _, section := range node.sections {
			// we have to use the external links, like ordinary author to
			// refer among files
			fmt.Printf("%s- [%s](#%s)\n",
				strings.Repeat("  ", section.level-1),
				section.title,
				filepath.Join(path, section.anchor))
		}

		fmt.Println("and subdirectories:")

		for _, d := range node.subdirs {
			fmt.Printf("- [directory %q](#%s)\n",
				d.dir,
				filepath.Join(path, d.dir))
		}

		defer wg.Done()
	}()

	for _, d := range node.subdirs {
		addIndex(out, d, path, wg)
	}
}
