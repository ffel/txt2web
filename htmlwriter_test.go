package txt2web

import "fmt"

func ExampleWriteHtml() {
	txtRoot = "example"
	destinations := WriteHtml(Generate(namegen("fileb.txt")))

	for d := range destinations {
		fmt.Printf("-- %s:\n%s\n", d.Path, d.Contents)
	}

	// output:
	// -- fileb.txt:
	// Pellentesque lobortis lacus
	// ===========================
	//
	// Condimentum rutrum enim blandit. Sed vitae luctus libero. Aliquam erat
	// volutpat. Morbi accumsan sem sodales lorem congue placerat. Nam auctor
	// sapien id libero vulputate, non sodales nibh tempus. Donec sagittis
	// consectetur magna sit amet vehicula. Vivamus sit amet dui eget urna
	// vestibulum gravida. Quisque et mauris vehicula, maximus nisi luctus,
	// pellentesque dui.
}
