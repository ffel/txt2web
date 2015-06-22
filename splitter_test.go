package txt2web

func ExampleSplit() {
	txtRoot = "example"
	chunkterm(Split(Generate(namegen("filea.txt", "fileb.txt", "dira/filec.txt", "dira/filed.txt", "dirb/filee.txt"))))

	// it is reference that inserts path names

	// output:
	// [Een Lorem ipsum dolor sit amet](#een-lorem-ipsum-dolor-sit-amet)
	// [Twee Morbi finibus rutrum condimentum.](#twee-morbi-finibus-rutrum-condimentum.)
	// [Drie Pellentesque lobortis lacus](#drie-pellentesque-lobortis-lacus)
	// [Vier Nulla euismod placerat nunc at mattis](#vier-nulla-euismod-placerat-nunc-at-mattis)
	// [Vijf Donec lacus leo](#vijf-donec-lacus-leo)
	// [Zes Fusce non aliquet tortor.](#zes-fusce-non-aliquet-tortor.)
	// [Zeven Nulla ut faucibus felis](#zeven-nulla-ut-faucibus-felis)
	// [Acht Pellentesque lacinia](#acht-pellentesque-lacinia)
	// [Negen Vivamus eget cursus erat, in pharetra neque](#negen-vivamus-eget-cursus-erat-in-pharetra-neque)
	// [Tien Phasellus lorem eros](#tien-phasellus-lorem-eros)
}
