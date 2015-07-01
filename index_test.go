package txt2web

// use the test set-up in refs_test - make this more flexible

var index_testinputs []string = []string{`
Section Root
=========
`, `
Section A
=========
`, `
Section A2
==========
`, `
Section AA
==========
`, `
Section C1
==========

Section C2
==========

Section C2.1
------------

Section C2.2
------------
`,
}

func ExampleIndex() {
	markdownTerm(Index(setFiles(contentGen(index_testinputs...), "f0.txt", "a/f1.txt", "a/f2.txt", "aa/f3.txt", "a/b/c/f4.txt")))

	// output:
	// foo
}
