package txt2web

func ExampleSplit() {
	txtRoot = "example"
	chunkterm(Split(Generate(namegen("filea.txt", "fileb.txt", "dira/filec.txt", "dira/filed.txt", "dirb/filee.txt"))))

	// output:
	// [Een Lorem ipsum dolor sit amet](#een-lorem-ipsum-dolor-sit-amet)
	// [Twee Morbi finibus rutrum condimentum.](#twee-morbi-finibus-rutrum-condimentum.)
	// [Images](#images)
	// [Drie Pellentesque lobortis lacus](#drie-pellentesque-lobortis-lacus)
	// [Vier Nulla euismod placerat nunc at mattis](#dira/vier-nulla-euismod-placerat-nunc-at-mattis)
	// [Vijf Donec lacus leo](#dira/vijf-donec-lacus-leo)
	// [Zes Fusce non aliquet tortor.](#dira/zes-fusce-non-aliquet-tortor.)
	// [Zeven Nulla ut faucibus felis](#dira/zeven-nulla-ut-faucibus-felis)
	// [Acht Pellentesque lacinia](#dirb/acht-pellentesque-lacinia)
	// [Negen Vivamus eget cursus erat, in pharetra neque](#dirb/negen-vivamus-eget-cursus-erat-in-pharetra-neque)
	// [Tien Phasellus lorem eros](#dirb/tien-phasellus-lorem-eros)
}
