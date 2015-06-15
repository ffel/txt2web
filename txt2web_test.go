package txt2web

import (
	"encoding/json"
	"fmt"
)

func Example() {
	outc := Convert("example/dira", "static")

	// something wrong in wrapper I guess ...

	for f := range outc {
		fmt.Println("%v -- %v\n", f.Path, string(f.Contents))
	}

	// output:
	// boo
}

func Example_1() {
	txtRoot = "example/dirb"
	destination = "static"
	outc := Split(Generate(TxtFiles(txtRoot, destination)))

	for c := range outc {
		bytes, err := json.Marshal(c.Json)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(bytes))
		}
	}

	// output:
	// baz
}
