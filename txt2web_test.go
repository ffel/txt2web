package txt2web

import "fmt"

func Example() {
	outc := Convert("example", "static")

	for f := range outc {
		fmt.Printf("%v -- %v\n", f.Path, string(f.Contents))
	}

	// output:
	// filea.txt -- <h1 id="morbi-finibus-rutrum-condimentum.">Morbi finibus rutrum condimentum.</h1>
	// <p>Duis posuere libero vitae auctor rhoncus. Donec eleifend malesuada ligula, vel condimentum ipsum tincidunt a. Nam eget ante eget tellus rhoncus auctor eu eu augue.</p>
	//
	// filea.txt -- <h1 id="lorem-ipsum-dolor-sit-amet">Lorem ipsum dolor sit amet</h1>
	// <p>Consectetur adipiscing elit. Sed fringilla mi sed sapien tempor, vel vulputate tortor vulputate. Integer a risus interdum, ornare magna pulvinar, molestie neque. Quisque risus tortor, pretium eu convallis eget, mattis vel tortor.</p>
	//
	// fileb.txt -- <h1 id="pellentesque-lobortis-lacus">Pellentesque lobortis lacus</h1>
	// <p>Condimentum rutrum enim blandit. Sed vitae luctus libero. Aliquam erat volutpat. Morbi accumsan sem sodales lorem congue placerat. Nam auctor sapien id libero vulputate, non sodales nibh tempus. Donec sagittis consectetur magna sit amet vehicula. Vivamus sit amet dui eget urna vestibulum gravida. Quisque et mauris vehicula, maximus nisi luctus, pellentesque dui.</p>
	//
	// dira/filec.txt -- <h1 id="donec-lacus-leo">Donec lacus leo</h1>
	// <p>Feugiat sit amet pulvinar eget, rutrum non nisi. Donec porttitor rutrum cursus. Nam vel interdum purus. Nunc lobortis maximus lectus eu consectetur. Vivamus ut fringilla justo. Suspendisse dictum dignissim suscipit.</p>
	//
	// dira/filec.txt -- <h1 id="nulla-euismod-placerat-nunc-at-mattis">Nulla euismod placerat nunc at mattis</h1>
	// <p>Pellentesque nisi ipsum, dapibus sit amet lacus et, ornare finibus orci. Suspendisse sit amet tristique arcu, eget rhoncus arcu. Suspendisse potenti. Sed eu dolor ut dui bibendum pretium et nec augue. Sed euismod est sit amet mi posuere, vel iaculis urna sodales.</p>
	//
	// dira/filed.txt -- <h1 id="nulla-ut-faucibus-felis">Nulla ut faucibus felis</h1>
	// <p>A pellentesque ligula. Pellentesque metus sapien, laoreet non rhoncus et, congue non velit. Integer elementum sagittis eros id suscipit.</p>
	//
	// dira/filed.txt -- <h1 id="fusce-non-aliquet-tortor.">Fusce non aliquet tortor.</h1>
	// <p>Pellentesque vitae odio in justo lacinia aliquet. Mauris suscipit urna eget odio facilisis varius a eu orci. Maecenas scelerisque quam ac venenatis ullamcorper. Duis scelerisque purus enim, vel iaculis velit convallis ac.</p>
	//
	// dirb/filee.txt -- <h1 id="phasellus-lorem-eros">Phasellus lorem eros</h1>
	// <p>commodo mollis consectetur ut, rhoncus imperdiet nisi. Pellentesque non elementum sapien. Morbi dignissim fringilla augue. Interdum et malesuada fames ac ante ipsum primis in faucibus. Sed porttitor at libero in mollis.</p>
	//
	// dirb/filee.txt -- <h1 id="vivamus-eget-cursus-erat-in-pharetra-neque">Vivamus eget cursus erat, in pharetra neque</h1>
	// <p>Aliquam quis lorem nisi. Aenean luctus eros at nunc porta, sit amet pretium ex pellentesque.</p>
	// <p>Vestibulum ut ultricies velit. Duis fermentum sit amet odio a efficitur. Nullam a finibus nisi, in porta nisl. Ut sed est sit amet nisi vehicula lacinia.</p>
	//
	// dirb/filee.txt -- <h1 id="pellentesque-lacinia">Pellentesque lacinia</h1>
	// <p>velit vel cursus convallis, sapien justo ultricies felis, a aliquet nunc tellus quis eros. Ut et rutrum turpis. Fusce risus magna, placerat eu lectus eget, feugiat malesuada nisl. Donec a laoreet ante, in laoreet velit.</p>
	// <h2 id="duis-faucibus-auctor-tortor-nec-accumsan">Duis faucibus auctor tortor nec accumsan</h2>
	// <p>Vivamus erat quam, ultrices in aliquam vitae, dapibus non lacus. Phasellus gravida ligula urna, at congue dui vulputate vitae. Phasellus massa velit, lacinia in tincidunt vel, pharetra id urna. Morbi vel massa malesuada justo euismod dapibus. Vivamus laoreet lacinia urna eget dapibus.</p>
	// <h2 id="vivamus-luctus-maximus-risus">Vivamus luctus maximus risus</h2>
	// <p>sed tincidunt nisi fermentum ac. Proin suscipit bibendum odio, eu elementum massa auctor imperdiet. Curabitur erat urna, eleifend nec viverra lacinia, volutpat porttitor dolor. Mauris a neque pharetra libero volutpat gravida. Mauris vitae eros urna. Nunc viverra consequat volutpat.</p>
}
