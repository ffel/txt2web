package txt2web

import "testing"

var images_testinputs []string = []string{
	`![](image.png)
`, `![](images/image.png)
`, `![](~/images/image.png)
`, `![](images/image.png)
`,
}

func TestImages(t *testing.T) {
	inout := []struct{ in, out string }{
		{images_testinputs[0], images_testinputs[1]},
		{images_testinputs[2], images_testinputs[3]},
	}

	for _, tt := range inout {
		c := markdownChan(Images(setFiles(contentGen(tt.in), "path/file.txt")))

		got := string(<-c)

		if got != tt.out {
			t.Errorf("---- expected:\n%q\n---- got:\n\n%q", tt.out, got)
		}
	}
}
