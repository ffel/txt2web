package txt2web

import (
	"fmt"

	"github.com/ffel/piperunner"
)

func init() {
	piperunner.StartPool()
}

// see
// https://github.com/ffel/pandocfilter/tree/master/cmd/pdtoc/pdtoc.go
// for some inspiration

func collectheaders(file string) []Header {
	fmt.Printf("file: %s\n", file)
	return make([]Header, 0)
}
