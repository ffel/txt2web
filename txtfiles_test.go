package txt2web

import "fmt"

func ExampleRun() {
	// this test fails if example/static exists ...
	runterm(TxtFiles("example", "example/static"))
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
