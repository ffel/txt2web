package txt2web

// txtfiles finds text files and works as a pre-generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// TxtFiles starts processing text files
// destination is used to make sure TxtFiles ignores destination
// in case it is a subdirectory of root.
func TxtFiles(root, dest string) <-chan string {
	txtRoot = root

	var err error
	destination, err = filepath.Abs(dest)

	if err != nil {
		log.Fatalf("destination problem: %v\n", err)
	}

	filenames := make(chan string)
	go func() {
		walk(".", filenames)
		close(filenames)
	}()
	return filenames
}

// walk is the recursive function to find text files in txtRoot
func walk(path string, filenames chan string) {
	pathname, err := filepath.Abs(filepath.Join(txtRoot, path))

	if err != nil {
		log.Fatalf("path name problem: %v\n", err)
	}

	if pathname == destination {
		fmt.Printf("ignore destination %q\n", path)
		return
	}

	files, err := ioutil.ReadDir(filepath.Join(txtRoot, path))

	if err != nil {
		log.Println("read dir problem:", err)
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
