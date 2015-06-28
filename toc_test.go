package txt2web

func ExampleToc() {
	txtRoot = "example"
	out := Toc(Generate(namegen("dirb/filee.txt")))

	for o := range out {
		_ = o
	}

	// output:
	// foo
}
