package main

import (
	"flag"
	"fmt"
	"io"
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
	images          = "images"
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

	htmlc := txt2web.Convert(src, dest, createFuncProcessImage(src, dest))

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

func createFuncProcessImage(sourceDir, targetDir string) txt2web.FuncProcessImage {

	return func(targetSource map[string]string) {
		for t, s := range targetSource {

			srcName := filepath.Join(sourceDir, s)
			trgName := filepath.Join(targetDir, t)

			fmt.Printf("copy %q to %q\n", srcName, trgName)

			err := os.MkdirAll(filepath.Dir(trgName), dirPermissions)
			if err != nil {
				log.Println("t2w image dir - cannot create dir.", err)
			}

			r, err := os.Open(srcName)
			if err != nil {
				log.Println("t2w image files - cannot read image.", err)
				return
			}
			defer r.Close()

			w, err := os.Create(trgName)
			if err != nil {
				log.Println("t2w image files - cannot create image.", err)
				return
			}
			defer w.Close()

			_, err = io.Copy(w, r)
			if err != nil {
				log.Println("t2w image files - cannot copy image.", err)
				return
			}
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
