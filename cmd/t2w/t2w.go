package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ffel/txt2web"
)

// run - go run t2w.go -src=../../example/ -dest=../../static

func main() {
	var src string
	var dest string

	flag.StringVar(&src, "src", ".", "source root of txt tree")
	flag.StringVar(&dest, "dest", "static", "destination root of html tree")

	flag.Parse()

	htmlc := txt2web.Convert(src)

	for h := range htmlc {
		// create deep directory
		err := os.MkdirAll(filepath.Join(dest, filepath.Dir(h.Path)), 0755)

		if err != nil {
			log.Println(err)
			continue
		}

		// swap .txt extension for .html
		path := strings.TrimSuffix(h.Path, filepath.Ext(h.Path)) + ".html"

		// write file
		err = ioutil.WriteFile(filepath.Join(dest, path), h.Contents, 0644)

		if err != nil {
			log.Println(err)
		}
	}
}
