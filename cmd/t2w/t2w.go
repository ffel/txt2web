package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ffel/txt2web"
)

// run - go run t2w.go -src=../../example/ -dest=../../example_html

// pages subdirectory
const pages = "pages"

func main() {
	var src string
	var dest string

	flag.StringVar(&src, "src", ".", "source root of txt tree")
	flag.StringVar(&dest, "dest", "static", "destination root of html tree")

	flag.Parse()

	htmlc := txt2web.Convert(src, dest)

	for h := range htmlc {

		// create deep directory
		err := os.MkdirAll(filepath.Join(dest, pages, filepath.Dir(h.Path)), 0755)

		if err != nil {
			log.Println(err)
			continue
		}

		// write file
		err = ioutil.WriteFile(filepath.Join(dest, pages, h.Path), h.Contents, 0644)

		if err != nil {
			log.Println(err)
		}
	}
}
