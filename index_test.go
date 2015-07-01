package txt2web

// use the test set-up in refs_test - make this more flexible

var index_testinputs []string = []string{`
Section A
=========
`, `
Section B
=========
`, `
Section C
=========
`,
}

func ExampleIndex() {
	markdownTerm(Index(setFiles(contentGen(index_testinputs...), "f0.txt", "a/f1.txt", "a/b/c/f2.txt")))

	// output:
	// foo
}
