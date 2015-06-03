package txt2web

// next by then challange: replace references in json (maybe this is the first challenge)
// what to do with this json

func Example() {
	Walk("./example", "./example")

	// output:
	// + [Lorem ipsum dolor sit amet](#filea.txt/lorem-ipsum-dolor-sit-amet)
	// + [Morbi finibus rutrum condimentum.](#filea.txt/morbi-finibus-rutrum-condimentum.)
	// + [Pellentesque lobortis lacus](#fileb.txt/pellentesque-lobortis-lacus)
	// + [Nulla euismod placerat nunc at mattis](#dira/filec.txt/nulla-euismod-placerat-nunc-at-mattis)
	// + [Donec lacus leo](#dira/filec.txt/donec-lacus-leo)
	// + [Fusce non aliquet tortor.](#dira/filed.txt/fusce-non-aliquet-tortor.)
	// + [Nulla ut faucibus felis](#dira/filed.txt/nulla-ut-faucibus-felis)
	// + [Pellentesque lacinia](#dirb/filee.txt/pellentesque-lacinia)
	//   + [Duis faucibus auctor tortor nec accumsan](#dirb/filee.txt/duis-faucibus-auctor-tortor-nec-accumsan)
	//   + [Vivamus luctus maximus risus](#dirb/filee.txt/vivamus-luctus-maximus-risus)
	// + [Vivamus eget cursus erat, in pharetra neque](#dirb/filee.txt/vivamus-eget-cursus-erat-in-pharetra-neque)
	// + [Phasellus lorem eros](#dirb/filee.txt/phasellus-lorem-eros)
}
