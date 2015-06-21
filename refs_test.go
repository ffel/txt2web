package txt2web

const hello = `# hello

world!
`

func ExampleReferences() {
	markdownTerm(contentGen(hello))

	// output:
	// hello
	// =====
	//
	// world!
	// ---
}
