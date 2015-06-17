package txt2web

import "fmt"

func ExampleGenerate() {
	txtRoot = "example"
	chunkterm(Generate(namegen("filea.txt", "dira/filec.txt")))

	// output:
	// [filea.txt](#filechunk)
	// [dira/filec.txt](#dira/filechunk)
}

// namegen is a simple generator for Generator testing
func namegen(names ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, n := range names {
			out <- n
		}
		close(out)
	}()
	return out
}

// chunkterm is a simple chunk terminator
func chunkterm(in <-chan Chunk) {
	for c := range in {
		fmt.Println(c)
	}
}
