package txt2web

// ref before id!!

var tests = []string{
	`
# section 1

# section 2

see [section 1](#section-1)
	`, `
# section 1

see [section 2](#section-2)

# section 2
	`, `
# section 1

see [section 2](#foo)

# section 2 {#foo}
	`,
}

func ExampleReferences() {
	markdownTerm(References(setFile(contentGen(tests...), "path/file.txt")))

	// output:
	// section 1 {#/path/section-1}
	// =========
	//
	// section 2 {#/path/section-2}
	// =========
	//
	// see [section 1](#/path/section-1)
	// ---
	// section 1 {#/path/section-1}
	// =========
	//
	// see [section 2](#/path/section-2)
	//
	// section 2 {#/path/section-2}
	// =========
	// ---
	// section 1 {#/path/section-1}
	// =========
	//
	// see [section 2](#/path/foo)
	//
	// section 2 {#/path/foo}
	// =========
	// ---

}
