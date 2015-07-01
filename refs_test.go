package txt2web

import "testing"

const in0 = `
# section 1

see [section 1](#section-1)
`

const out0 = `section 1
=========

see [section 1](#/path/section-1)
`

const in1 = `
# section 1

# section 2

see [section 1](#section-1)
`

const out1 = `section 1
=========

section 2
=========

see [section 1](#/path/section-1)
`

const in2 = `
# section 1

see [section 2](#section-2)

# section 2
`

const out2 = `section 1
=========

see [section 2](#/path/section-2)

section 2
=========
`

const in3 = `
# section 1

see [section 2](#foo)

# section 2 {#foo}
`

const out3 = `section 1
=========

see [section 2](#/path/foo)

section 2 {#foo}
=========
`

func TestReferences(t *testing.T) {
	inout := []struct{ in, out string }{
		{in0, out0},
		{in1, out1},
		{in2, out2},
		{in3, out3},
	}

	for _, tt := range inout {
		c := markdownChan(References(setFiles(contentGen(tt.in), "path/file.txt")))

		got := string(<-c)

		if got != tt.out {
			t.Errorf("---- expected:\n%q\n---- got:\n\n%q", tt.out, got)
		}
	}
}
