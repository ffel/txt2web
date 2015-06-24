package txt2web

import (
	"fmt"
	"sort"
)

func Example() {
	outc := Convert("example", "static")

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
	// dirb/acht-pellentesque-lacinia.html -- Acht Pellentesque lacinia
	// drie-pellentesque-lobortis-lacus.html -- Drie Pellentesque lobortis lacus
	// een-lorem-ipsum-dolor-sit-amet.html -- Een Lorem ipsum dolor sit amet
	// dirb/negen-vivamus-eget-cursus-erat-in-pharetra-neque.html -- Negen Vivamus eget cursus erat, in pharetra neque
	// dirb/tien-phasellus-lorem-eros.html -- Tien Phasellus lorem eros
	// twee-morbi-finibus-rutrum-condimentum..html -- Twee Morbi finibus rutrum condimentum.
	// dira/vier-nulla-euismod-placerat-nunc-at-mattis.html -- Vier Nulla euismod placerat nunc at mattis
	// dira/vijf-donec-lacus-leo.html -- Vijf Donec lacus leo
	// dira/zes-fusce-non-aliquet-tortor..html -- Zes Fusce non aliquet tortor.
	// dira/zeven-nulla-ut-faucibus-felis.html -- Zeven Nulla ut faucibus felis
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
