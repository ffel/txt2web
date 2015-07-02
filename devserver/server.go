package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

/*
This is a very simple test server to host the example web site.

$ cd txt2web			# sic, intended to run from the main package
$ go run devserver/server.go
*/

const path = "./example_html"

func main() {
	fmt.Println("Serve files from " + path)

	url := "http://localhost:8000/#"

	// start a browser
	go func() {
		time.Sleep(100 * time.Millisecond)

		// http://stackoverflow.com/questions/10377243
		var err error
		switch runtime.GOOS {
		case "linux":
			// untested
			err = exec.Command("xdg-open", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		case "windows":
			// untested
			exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}

		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir(path))))
}
