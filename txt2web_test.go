package txt2web

// next challange: path names should go without argument path, that is, relative to
// argument
// see - http://godoc.org/path/filepath#example-Rel
//
// besides, use http://godoc.org/path/filepath#Join

func Example() {
	Walk("./example", "./example")

	// output:
	// * Lorem ipsum dolor sit amet  (lorem-ipsum-dolor-sit-amet, filea.txt):
	//
	// * Morbi finibus rutrum condimentum.  (morbi-finibus-rutrum-condimentum., filea.txt):
	//
	// * Pellentesque lobortis lacus  (pellentesque-lobortis-lacus, fileb.txt):
	//
	// * Nulla euismod placerat nunc at mattis  (nulla-euismod-placerat-nunc-at-mattis, dira/filec.txt):
	//
	// * Donec lacus leo  (donec-lacus-leo, dira/filec.txt):
	//
	// * Fusce non aliquet tortor.  (fusce-non-aliquet-tortor., dira/filed.txt):
	//
	// * Nulla ut faucibus felis  (nulla-ut-faucibus-felis, dira/filed.txt):
	//
	// * Pellentesque lacinia  (pellentesque-lacinia, dirb/filee.txt):
	//
	// * Duis faucibus auctor tortor nec accumsan  (duis-faucibus-auctor-tortor-nec-accumsan, dirb/filee.txt):
	//
	// * Vivamus luctus maximus risus  (vivamus-luctus-maximus-risus, dirb/filee.txt):
	//
	// * Vivamus eget cursus erat, in pharetra neque  (vivamus-eget-cursus-erat-in-pharetra-neque, dirb/filee.txt):
	//
	// * Phasellus lorem eros  (phasellus-lorem-eros, dirb/filee.txt):
}
