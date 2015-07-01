package txt2web

import "testing"

var ref_testcase []string = []string{`
# section 1

see [section 1](#section-1)
`, `section 1
=========

see [section 1](#/path/section-1)
`, /* case two */ `
# section 1

# section 2

see [section 1](#section-1)
`, `section 1
=========

section 2
=========

see [section 1](#/path/section-1)
`, /* case three */ `
# section 1

see [section 2](#section-2)

# section 2
`, `section 1
=========

see [section 2](#/path/section-2)

section 2
=========
`, /* case four */ `
# section 1

see [section 2](#foo)

# section 2 {#foo}
`, `section 1
=========

see [section 2](#/path/foo)

section 2 {#foo}
=========
`}

func TestReferences(t *testing.T) {
	inout := []struct{ in, out string }{
		{ref_testcase[0], ref_testcase[1]},
		{ref_testcase[2], ref_testcase[3]},
		{ref_testcase[4], ref_testcase[5]},
		{ref_testcase[6], ref_testcase[7]},
	}

	for _, tt := range inout {
		c := markdownChan(References(setFiles(contentGen(tt.in), "path/file.txt")))

		got := string(<-c)

		if got != tt.out {
			t.Errorf("---- expected:\n%q\n---- got:\n\n%q", tt.out, got)
		}
	}
}
