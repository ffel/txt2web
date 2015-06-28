package main

import (
	"flag"
	"os"

	"github.com/ffel/txt2web"
)

// run - go run t2toc.go -src=../../example/ -dest=../../example_html

func main() {
	var src string
	var dest string

	flag.StringVar(&src, "src", ".", "source root of txt tree")
	flag.StringVar(&dest, "dest", "static", "destination root of html tree")

	flag.Parse()

	outc := txt2web.Toc(txt2web.Generate(txt2web.TxtFiles(src, dest)), os.Stdout)

	// we need to read outc to prevent blocking

	for _ = range outc {

	}
}
