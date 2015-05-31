package txt2web

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Header struct {
	Header string
	Key    string
	Path   string
}

// String prints Header
func (h Header) String() string {
	return fmt.Sprintf("%s (%s, %s)", h.Header, h.Key, h.Path)
}

// Walk start recursive iteration over sub dir tree
func Walk(path string) {
	for _, h := range Headers(path) {
		c, err := Contents(h)

		if err != nil {
			fmt.Printf("error %v\n", err)
		} else {
			fmt.Printf("%v:\n%s\n", h, c)
		}
	}

	for _, s := range SubDirs(path) {
		Walk(s)
	}
}

// Get all headers in path
func Headers(path string) []Header {
	files, err := ioutil.ReadDir(path)

	result := make([]Header, 0)

	if err != nil {
		return result
	}

	for _, f := range files {
		if f.Mode().IsRegular() && filepath.Ext(f.Name()) == ".txt" {
			result = append(result, collectheaders(path+"/"+f.Name())...)
		}
	}

	return result
}

// Get the contents of header
func Contents(header Header) (string, error) {
	return "", nil
}

// Get all sub directories in path such that these can be used in Headers()
func SubDirs(path string) []string {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return make([]string, 0)
	}

	result := make([]string, 0, len(files))

	for _, f := range files {
		if f.IsDir() {
			result = append(result, path+"/"+f.Name())
		}
	}

	return result
}
