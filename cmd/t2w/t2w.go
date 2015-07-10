package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ffel/txt2web"
)

// run - go run t2w.go -src=../../example/ -dest=../../example_html -server

// pages subdirectory
const (
	pages           = "pages"
	dirPermissions  = 0755
	filePermissions = 0644
)

func main() {
	var src string
	var dest string

	flag.StringVar(&src, "src", ".", "source root of txt tree")
	flag.StringVar(&dest, "dest", "static", "destination root of html tree")
	serverPtr := flag.Bool("server", false, "add simple file server")

	flag.Parse()

	if *serverPtr {
		if err := os.MkdirAll(dest, dirPermissions); err != nil {
			log.Println("Error creating server dir:", err)
		}

		if err := ioutil.WriteFile(filepath.Join(dest, "server.go"),
			[]byte(server), filePermissions); err != nil {
			log.Println("Error writing server:", err)
		}
	}

	htmlc := txt2web.Convert(src, dest)

	for h := range htmlc {

		// create deep directory
		err := os.MkdirAll(filepath.Join(dest, pages, filepath.Dir(h.Path)), dirPermissions)

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

const server = `package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

const path = "."
const port = ":4550"

func main() {
	url := "http://localhost" + port + "/#"

	go func() {
		var err error

		switch runtime.GOOS {
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}

		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(path))))
}
`
