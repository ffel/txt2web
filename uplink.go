package txt2web

// UpLinkNode takes section headers and translates these into links
// to the containing sections.
func UpLinkNode(in <-chan Chunk) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		for c := range in {
			out <- c
		}
		close(out)
	}()

	return out
}
