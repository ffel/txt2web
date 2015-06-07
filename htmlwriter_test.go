package txt2web

import "time"

func ExampleWriteHtml() {
	txtRoot = "example"
	WriteHtml(Generate(namegen("fileb.txt")))

	// there is a slight problem with this test.  WriteHtml uses
	// goroutines to print, but, WriteHtml is already terminated by then.
	// This makes that another test receives the output.

	// give the goroutives some time to terminate
	time.Sleep(time.Second)

	// output:
	// fileb.txt:
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
