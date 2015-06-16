package txt2web

import "fmt"

func Example() {
	outc := Convert("example", "static")

	for f := range outc {
		fmt.Printf("%v -- %v\n", f.Path, f.Title)
	}

	// output:
	// filea.txt -- Lorem ipsum dolor sit amet
	// filea.txt -- Morbi finibus rutrum condimentum.
	// fileb.txt -- Pellentesque lobortis lacus
	// dira/filec.txt -- Nulla euismod placerat nunc at mattis
	// dira/filec.txt -- Donec lacus leo
	// dira/filed.txt -- Fusce non aliquet tortor.
	// dira/filed.txt -- Nulla ut faucibus felis
	// dirb/filee.txt -- Pellentesque lacinia
	// dirb/filee.txt -- Vivamus eget cursus erat, in pharetra neque
	// dirb/filee.txt -- Phasellus lorem eros
}
