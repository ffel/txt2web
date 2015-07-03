package main

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
