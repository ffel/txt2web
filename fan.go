package txt2web

import "sync"

// fan-in and fan-out (duplicate, more than distribute)

// FuncC2H is a function that takes a Chunk channel and produces on an HtmlFile
// channel
type FuncC2H func(<-chan Chunk) <-chan HtmlFile

// MultiplexC2H duplicates messages on "in" over functions in "over"
// see http://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters-in-go
func MultiplexC2H(in <-chan Chunk, over ...FuncC2H) []<-chan HtmlFile {

	// out is the collection of over return values
	out := make([]<-chan HtmlFile, len(over))

	// fan will be the inbound channel which distributes in
	fan := make([]chan Chunk, len(over))

	// initialise fan, set relation between fan[i] and out[i]
	for i := 0; i < len(over); i++ {
		fan[i] = make(chan Chunk)
		out[i] = over[i](fan[i])
	}

	// wait for messages on in, distribute over fan
	go func() {
		for cin := range in {
			for _, cout := range fan {
				cout <- cin
			}
		}
		for _, f := range fan {
			close(f)
		}
	}()

	return out
}

// MergeH2H takes several channels and combine their input
// taken from http://blog.golang.org/pipelines, fan-in, fan-out
func MergeH2H(cs ...<-chan HtmlFile) <-chan HtmlFile {
	var wg sync.WaitGroup
	out := make(chan HtmlFile)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan HtmlFile) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
