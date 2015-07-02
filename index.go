package txt2web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/ffel/pandocfilter"
	"github.com/ffel/piperunner"
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
		addIndex(out, root, "", &wg)
		wg.Wait()

		close(out)
	}()

	return out
}

func addIndex(out chan Chunk, node *indexInfo, path string, wg *sync.WaitGroup) {

	wg.Add(1)

	path = filepath.Join(path, node.dir)

	go func() {
		defer wg.Done()

		sections := ""

		for _, section := range node.sections {
			// we have to use the external links, like ordinary author to
			// refer among files
			if section.level == 1 {
				sections += fmt.Sprintf("%s-   [%s](#%s)\n",
					strings.Repeat("    ", section.level-1),
					section.title,
					filepath.Join(path, section.anchor))
			} else {
				sections += fmt.Sprintf("%s-   %s\n",
					strings.Repeat("    ", section.level-1),
					section.title)
			}
		}

		directories := ""

		for _, d := range node.subdirs {
			directories += fmt.Sprintf("- [directory %q](#%s/index)\n",
				d.dir,
				filepath.Join(path, d.dir))
		}

		header := ""

		pathelem := strings.Split(path, "/")

		if path != "." && len(pathelem) >= 1 {
			header = fmt.Sprintf("[Index](#%v)", filepath.Join(pathelem[:len(pathelem)-1]...))
		} else {
			header = "Index"
		}

		t := template.New("index")
		t, err := t.Parse(indexTxt)
		if err != nil {
			log.Fatal(err)
		}

		buff := bytes.NewBufferString("")

		err = t.Execute(buff, struct{ Index, Sections, Directories string }{header, sections, directories})
		if err != nil {
			log.Fatal(err)
		}

		resultc := piperunner.Exec("pandoc -f markdown -t json", buff.Bytes())

		result := <-resultc

		if err := result.Err; err != nil {
			log.Fatal(result.Text)
		}

		var jsondata interface{}
		err = json.Unmarshal(result.Text, &jsondata)

		if err != nil {
			log.Fatal(err)
		}

		out <- Chunk{Json: jsondata, Path: filepath.Join(path, "index.txt"), PandocId: ""}

	}()

	for _, d := range node.subdirs {
		addIndex(out, d, path, wg)
	}
}

// pandoc reformats this markdown, so the final result is different
const indexTxt = `
# {{.Index}}

## Sections

{{.Sections}}

## Directories

{{.Directories}}
`
