package txt2web

// The references is the node in the pipeline that replaces markdown references
// (which work within markdown files) for references that work between html
// pages.
//
// This node is a "two subprocess" node.  The first sub-process finds the
// id's, the second sub-process translates references to those id's

type RefChunk struct {
	Chunk                          // original chunk
	Translations map[string]string // internal ref -> external ref
}

func References(in <-chan Chunk) <-chan Chunk {
	return ref_translator(ref_finder(in))
}

func ref_finder(in <-chan Chunk) <-chan RefChunk {
	out := make(chan RefChunk)
	go func() {
		for c := range in {
			out <- RefChunk{c, make(map[string]string)}
		}
		close(out)
	}()
	return out
}

func ref_translator(in <-chan RefChunk) <-chan Chunk {
	out := make(chan Chunk)
	go func() {
		for rc := range in {
			out <- rc.Chunk
		}
		close(out)
	}()
	return out
}
