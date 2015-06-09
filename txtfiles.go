package txt2web

// txtfiles finds text files and works as a pre-generator

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// any chance there are two instances running at the same time?
// not for now, so no need to store path to web root and target
// path with every file

// TxtFiles starts processing text files
func TxtFiles(root string) <-chan string {
	txtRoot = root
	filenames := make(chan string)
	go func() {
		walk(".", filenames)
		close(filenames)
	}()
	return filenames
}

// walk is the recursive function to find text files in txtRoot
func walk(path string, filenames chan string) {
	files, err := ioutil.ReadDir(filepath.Join(txtRoot, path))

	if err != nil {
		log.Println(err)
		return
	}

	// files first
	for _, f := range files {
		if f.Mode().IsRegular() && filepath.Ext(f.Name()) == ".txt" {
			filenames <- filepath.Join(path, f.Name())
		}
	}

	// dive into dirs
	for _, d := range files {
		if d.IsDir() {
			walk(d.Name(), filenames)
		}
	}
}
