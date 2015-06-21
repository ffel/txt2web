package txt2web

// ref before id!!

var tests = []string{
	`
# section 1

# section 2

see [section 1](#section-1)
	`,
}

func ExampleReferences() {
	markdownTerm(References(contentGen(tests...)))

	// output:
	// section 1
	// =========
	//
	// section 2
	// =========
	//
	// see [section 1](foobar)
	// ---
}
