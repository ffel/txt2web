package txt2web

import "testing"

var uplink_tests []string = []string{`
main
====

sub 1
-----

sub 2
-----
`, `
boo
`,
}

func TestUpLinks(t *testing.T) {
	inout := []struct{ in, out string }{
		{uplink_tests[0], uplink_tests[1]},
	}

	for _, tt := range inout {
		c := markdownChan(UpLinkNode(setFiles(contentGen(tt.in), "file.txt")))

		got := string(<-c)

		if got != tt.out {
			t.Errorf("---- expected:\n%q\n---- got:\n\n%q", tt.out, got)
		}
	}
}
