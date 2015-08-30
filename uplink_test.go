package txt2web

import "testing"

var uplink_tests []string = []string{`
main
====

sub A
-----

### Sub A.A

sub B
-----

### Sub B.A

### Sub B.B

main II
=======
`, `
foo
`, `
### weird a

# weird b

###### weird c

###### weird c2
`, `
bar
`,
}

// for one test only
// go test -run TestUpLinks

func TestUpLinks(t *testing.T) {
	inout := []struct{ in, out string }{
		{uplink_tests[0], uplink_tests[1]},
		{uplink_tests[2], uplink_tests[3]},
	}

	for _, tt := range inout {
		c := markdownChan(UpLinkNode(setFiles(contentGen(tt.in), "file.txt")))

		got := string(<-c)

		if got != tt.out {
			t.Errorf("---- expected:\n%q\n---- got:\n\n%q", tt.out, got)
		}
	}
}
