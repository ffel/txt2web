package txt2web

import (
	"fmt"
	"sort"
)

func Example() {
	outc := Convert("example/dirb", "static")

	// the order in which results appear in outc is not deterministic
	// as it depends on the speed the worker pool processes each chunk
	// thats why we sort the results

	var results []HtmlFile

	for f := range outc {
		results = append(results, f)
	}

	sort.Sort(byTitle(results))

	for _, r := range results {
		fmt.Printf("%v -- %v\n", r.Path, r.Title)
	}

	// output:
	// filee.txt -- Acht Pellentesque lacinia
	// filee.txt -- Negen Vivamus eget cursus erat, in pharetra neque
	// filee.txt -- Tien Phasellus lorem eros
}

type byTitle []HtmlFile

func (s byTitle) Len() int {
	return len(s)
}
func (s byTitle) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byTitle) Less(i, j int) bool {
	return s[i].Title < s[j].Title
}
