package txt2web

func ExampleSplit() {
	txtRoot = "example"
	chunkterm(Split(Generate(namegen("filea.txt", "fileb.txt", "dira/filec.txt", "dira/filed.txt", "dirb/filee.txt"))))

	// output:
	// [Een Lorem ipsum dolor sit amet](#filea.txt/een-lorem-ipsum-dolor-sit-amet)
	// [Twee Morbi finibus rutrum condimentum.](#filea.txt/twee-morbi-finibus-rutrum-condimentum.)
	// [Drie Pellentesque lobortis lacus](#fileb.txt/drie-pellentesque-lobortis-lacus)
	// [Vier Nulla euismod placerat nunc at mattis](#dira/filec.txt/vier-nulla-euismod-placerat-nunc-at-mattis)
	// [Vijf Donec lacus leo](#dira/filec.txt/vijf-donec-lacus-leo)
	// [Zes Fusce non aliquet tortor.](#dira/filed.txt/zes-fusce-non-aliquet-tortor.)
	// [Zeven Nulla ut faucibus felis](#dira/filed.txt/zeven-nulla-ut-faucibus-felis)
	// [Acht Pellentesque lacinia](#dirb/filee.txt/acht-pellentesque-lacinia)
	// [Negen Vivamus eget cursus erat, in pharetra neque](#dirb/filee.txt/negen-vivamus-eget-cursus-erat-in-pharetra-neque)
	// [Tien Phasellus lorem eros](#dirb/filee.txt/tien-phasellus-lorem-eros)
}
