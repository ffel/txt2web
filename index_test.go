package txt2web

func ExampleIndex() {
	txtRoot = "example"
	chunkterm(Index(Generate(namegen("filea.txt", "dira/filec.txt"))))

	// output:
	// [filea.txt](#filechunk)
	// [index.txt](#/)
	// [dira/filec.txt](#dira/filechunk)
	// [dirs/index.txt](#dira/)
}
