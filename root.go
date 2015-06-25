package txt2web

// The node `WriteRoot` adds the files in the root of the web site.

func WriteRoot(in <-chan Chunk) <-chan HtmlFile {
	out := make(chan HtmlFile)

	go func() {
		// see htmlwriter for more sophistication

		for c := range in {
			// ignore for now
			_ = c
		}
		close(out)
	}()

	return out
}
