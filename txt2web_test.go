package txt2web

func Example() {
	Walk("./example")

	// output:
	// file: ./example/filea.txt
	// file: ./example/fileb.txt
	// file: ./example/dira/filec.txt
	// file: ./example/dira/filed.txt
	// file: ./example/dirb/filee.txt
}
