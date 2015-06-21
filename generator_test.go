package txt2web

func ExampleGenerate() {
	txtRoot = "example"
	chunkterm(Generate(namegen("filea.txt", "dira/filec.txt")))

	// output:
	// [filea.txt](#filechunk)
	// [dira/filec.txt](#dira/filechunk)
}
