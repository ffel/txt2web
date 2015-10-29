package txt2web

import "testing"

/*
what's next .... (after a long period of silence)

this is handy: pythia -http=:8081 github.com/ffel/txt2web
*/

var uplink_tests []string = []string{`
# Hello
`, `[Hello](index.html)
===================
`, `
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
`, `[main](index.html)
==================

[sub A](main)
-------------

### [Sub A.A](sub-a)

[sub B](main)
-------------

### [Sub B.A](sub-b)

### [Sub B.B](sub-b)

[main II](index.html)
=====================
`, `
### weird a

# weird b

###### weird c

###### weird c2

### weird d
`, `### [weird a](index.html)

[weird b](index.html)
=====================

###### [weird c](weird-b)

###### [weird c2](weird-b)

### [weird d](weird-b)
`, `
# h1
# [h2](#h1)
`, `
baz
`,
}

// for one test only
// go test -run TestUpLinks

func TestUpLinks(t *testing.T) {
	inout := []struct{ in, out string }{
		{uplink_tests[0], uplink_tests[1]},
		{uplink_tests[2], uplink_tests[3]},
		{uplink_tests[4], uplink_tests[5]},
		// {uplink_tests[6], uplink_tests[7]}, // bedoeld om verschil in structuur te vinden
	}

	// println(jsonPrint(`[{"unMeta":{}},[{"t":"Header","c":[1,["hello",[],[]],[{"t":"Link","c":[[{"t":"Str","c":"Hello"}],["index.html",""]]}]]}]]`))

	for _, tt := range inout {
		c := markdownChan(structPrint(UpLinkNode(setFiles(contentGen(tt.in), "file.txt")), "**", false))

		got := string(<-c)

		if got != tt.out {
			t.Errorf("---- expected:\n%q\n---- got:\n\n%q", tt.out, got)
		}
	}
}
