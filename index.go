package txt2web

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Index is the node that analyzes chunks for syblings and childs and will
// add index chunks to the stream

// indexInfo is a node in a tree
type indexInfo struct {
	path string
	// sections []tocEntry
	subdirs []*indexInfo
}

func (ii *indexInfo) knowsDir(path string) (bool, *indexInfo) {
	for _, dir := range ii.subdirs {
		if dir.path == path {
			return true, dir
		}
	}
	return false, &indexInfo{}
}

func (ii *indexInfo) String() string {
	result := ii.path + ": [ "

	for _, c := range ii.subdirs {
		result += c.String() + ", "
	}

	return result + " ]"
}

// Index analyzes the in stream and adds index chunks to the stream
func Index(in <-chan Chunk) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		root := &indexInfo{path: "."}

		for c := range in {

			path := filepath.Dir(c.Path)

			current := root

			// if we're not in root, we have to determine current as a node
			// in the tree
			if path != "." {

				elem := strings.Split(path, "/")

				// walk the tree and create nodes where there are no nodes
				for _, dir := range elem {

					if ok, sub := current.knowsDir(dir); ok {
						current = sub
					} else {
						current.subdirs = append(current.subdirs, &indexInfo{path: dir})
						current = current.subdirs[len(current.subdirs)-1]
					}
				}
			}

			fmt.Printf("tree    %v\n", root)
			fmt.Printf("current %v\n", current)

			out <- c
		}
		close(out)
	}()

	return out
}

// we will use type tocEntry to collect sections in one file
