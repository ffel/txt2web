package txt2web

import "fmt"

func Example() {
	outc := Convert("example/dira", "static")

	for f := range outc {
		fmt.Printf("%v -- %v\n", f.Path, string(f.Contents))
	}

	// output:
	// filec.txt -- <h1 id="donec-lacus-leo">Donec lacus leo</h1>
	// <p>Feugiat sit amet pulvinar eget, rutrum non nisi. Donec porttitor rutrum cursus. Nam vel interdum purus. Nunc lobortis maximus lectus eu consectetur. Vivamus ut fringilla justo. Suspendisse dictum dignissim suscipit.</p>
	//
	// filec.txt -- <h1 id="nulla-euismod-placerat-nunc-at-mattis">Nulla euismod placerat nunc at mattis</h1>
	// <p>Pellentesque nisi ipsum, dapibus sit amet lacus et, ornare finibus orci. Suspendisse sit amet tristique arcu, eget rhoncus arcu. Suspendisse potenti. Sed eu dolor ut dui bibendum pretium et nec augue. Sed euismod est sit amet mi posuere, vel iaculis urna sodales.</p>
	//
	// filed.txt -- <h1 id="fusce-non-aliquet-tortor.">Fusce non aliquet tortor.</h1>
	// <p>Pellentesque vitae odio in justo lacinia aliquet. Mauris suscipit urna eget odio facilisis varius a eu orci. Maecenas scelerisque quam ac venenatis ullamcorper. Duis scelerisque purus enim, vel iaculis velit convallis ac.</p>
	//
	// filed.txt -- <h1 id="nulla-ut-faucibus-felis">Nulla ut faucibus felis</h1>
	// <p>A pellentesque ligula. Pellentesque metus sapien, laoreet non rhoncus et, congue non velit. Integer elementum sagittis eros id suscipit.</p>
	//
}
