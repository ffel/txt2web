package txt2web

import "fmt"

func ExampleRun() {
	runterm(TxtFiles("example"))
	// output:
	// filea.txt
	// fileb.txt
	// dira/filec.txt
	// dira/filed.txt
	// dirb/filee.txt
}

func runterm(in <-chan string) {
	for f := range in {
		fmt.Println(f)
	}
}
