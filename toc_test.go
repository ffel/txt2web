package txt2web

import "os"

func ExampleToc() {
	txtRoot = "example"
	out := Toc(Generate(namegen("dirb/filee.txt")), os.Stdout)

	for o := range out {
		_ = o
	}

	// output:
	// dirb/filee.txt:
	//                                                  #dirb/acht-pellentesque-lacinia
	// - [Acht Pellentesque lacinia](#acht-pellentesque-lacinia)
	//   - [Duis faucibus auctor tortor nec accumsan](#duis-faucibus-auctor-tortor-nec-accumsan)
	//   - [Vivamus luctus maximus risus](#vivamus-luctus-maximus-risus)
	//                           #dirb/negen-vivamus-eget-cursus-erat-in-pharetra-neque
	// - [Negen Vivamus eget cursus erat, in pharetra neque](#negen-vivamus-eget-cursus-erat-in-pharetra-neque)
	//                                                  #dirb/tien-phasellus-lorem-eros
	// - [Tien Phasellus lorem eros](#tien-phasellus-lorem-eros)
}
