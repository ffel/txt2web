package txt2web

// Index is the node that analyzes chunks for syblings and childs and will
// add index chunks to the stream

// Index analyzes the in stream and adds index chunks to the stream
func Index(in <-chan Chunk) <-chan Chunk {
	out := make(chan Chunk)

	go func() {
		for c := range in {

			// expect several file chunks with several sections per directory

			// identify subdirs first

			out <- c
		}
		close(out)
	}()

	return out
}

// we will use type tocEntry to collect sections in one file
