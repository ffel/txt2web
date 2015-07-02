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
	markdownTerm(Index(setFiles(contentGen(index_testinputs...), "f0.txt", "a/f1.txt", "a/f2.txt", "aa/f3.txt", "a/b/c/f4.txt")), true)

	// output:
	// <<0 - "a/b/c/f4.txt">>
	// Section C1
	// ==========

	// Section C2
	// ==========

	// Section C2.1
	// ------------

	// Section C2.2
	// ------------
	// <<1 - "a/b/c/index.txt">>
	// [Index](#a/b)
	// =============

	// Sections
	// --------

	// -   [Section C1](#a/b/c/section-c1)
	// -   [Section C2](#a/b/c/section-c2)
	//     -   Section C2.1
	//     -   Section C2.2

	// Directories
	// -----------
	// <<2 - "a/b/index.txt">>
	// [Index](#a)
	// ===========

	// Sections
	// --------

	// Directories
	// -----------

	// -   [directory "c"](#a/b/c)

	// <<3 - "a/f1.txt">>
	// Section A
	// =========
	// <<4 - "a/f2.txt">>
	// Section A2
	// ==========
	// <<5 - "a/index.txt">>
	// Index
	// =====

	// Sections
	// --------

	// -   [Section A](#a/section-a)
	// -   [Section A2](#a/section-a2)

	// Directories
	// -----------

	// -   [directory "b"](#a/b)

	// <<6 - "aa/f3.txt">>
	// Section AA
	// ==========
	// <<7 - "aa/index.txt">>
	// Index
	// =====

	// Sections
	// --------

	// -   [Section AA](#aa/section-aa)

	// Directories
	// -----------
	// <<8 - "f0.txt">>
	// Section Root
	// ============
	// <<9 - "index.txt">>
	// Index
	// =====

	// Sections
	// --------

	// -   [Section Root](#section-root)

	// Directories
	// -----------

	// -   [directory "a"](#a)
	// -   [directory "aa"](#aa)
}
