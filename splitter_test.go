package txt2web

func ExampleSplit() {
	txtRoot = "example"
	chunkterm(Split(Generate(namegen("filea.txt", "dira/filec.txt"))))

	// output:
	// [Lorem ipsum dolor sit amet](#filea.txt/lorem-ipsum-dolor-sit-amet)
	// [Morbi finibus rutrum condimentum.](#filea.txt/morbi-finibus-rutrum-condimentum.)
	// [Nulla euismod placerat nunc at mattis](#dira/filec.txt/nulla-euismod-placerat-nunc-at-mattis)
	// [Donec lacus leo](#dira/filec.txt/donec-lacus-leo)
}
