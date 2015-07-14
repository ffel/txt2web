package txt2web

import "testing"

// ignore 0
// each test creates a new instance of Images which resets the counter
// this package does not determine the final target path
var images_testinputs []string = []string{`
`, `![](image.png)
`, `![](images/image_1_1.png)
`, `![](image.png)

![](a/image.png)

![](b/c/image.png)
`, `![](images/image_1_1.png)

![](images/image_1_2.png)

![](images/image_1_3.png)
`, `![](~/images/image.png)
`, `![](images/image_1_1.png)
`, `![](http://example.com/image.png)

![](https://example.org/image.png)
`, `![](http://example.com/image.png)

![](https://example.org/image.png)
`,
}

func TestImages(t *testing.T) {
	inout := []struct{ in, out string }{
		{images_testinputs[1], images_testinputs[2]},
		{images_testinputs[3], images_testinputs[4]},
		{images_testinputs[5], images_testinputs[6]},
		{images_testinputs[7], images_testinputs[8]},
	}

	for _, tt := range inout {
		c := markdownChan(ImagesNode(setFiles(contentGen(tt.in), "path/file.txt"), "images", copyIgnore))

		got := string(<-c)

		if got != tt.out {
			t.Errorf("---- expected:\n%q\n---- got:\n\n%q", tt.out, got)
		}
	}
}
