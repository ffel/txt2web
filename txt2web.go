package txt2web

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/ffel/piperunner"
)

// Convert does the complete txt2web conversion
func Convert(txtroot, destination string) <-chan HtmlFile {

	var filenamec <-chan string

	// find file names, ignore possible clashes if destination is a sub map
	filenamec = TxtFiles(txtroot, destination)

	var chunkc <-chan Chunk

	// read entire files in chunks
	chunkc = Generate(filenamec)
	// replace anchors for angular routes
	chunkc = References(chunkc)
	// split chunks into one section per chunk
	chunkc = Split(chunkc)

	// the result of split is needed in two nodes
	var root chan Chunk = make(chan Chunk)
	var pages chan Chunk = make(chan Chunk)

	// this is not really a fan-out, chunks are not distributed, chunks are
	// duplicated
	go func(in <-chan Chunk) {
		for c := range in {
			root <- c
			pages <- c
		}
		close(root)
		close(pages)
	}(chunkc)

	var htmlc <-chan HtmlFile

	htmlc = MergeHtmlFileCh(WriteRoot(root), WriteHtml(pages))

	return htmlc
}

// Chunk is the basis data object for one #-section
type Chunk struct {
	Json     interface{} // internal representation of Json
	Path     string      // file + path local to web root
	Section  int         // nr of # section in file (0 is pre section text)
	Title    string      // section title
	PandocId string      // original key (either user or pandoc provided)
}

// Webkey is the chunk id that is used to refer between txt files
func (c Chunk) Webkey() string {
	return fmt.Sprintf("#%s", filepath.Join(filepath.Dir(c.Path), c.PandocId))
}

// String produces the markdown link for Chunk
func (c Chunk) String() string {
	return fmt.Sprintf("[%s](%s)", c.Title, c.Webkey())
}

// any chance there are two instances running at the same time?
// not for now, so no need to store path to web root and target
// path with every file

// txtRoot is relative or absolute path to root of txt2web project
var txtRoot string
var destination string

// HtmlFile is the contents and the file to write
type HtmlFile struct {
	Contents []byte
	Title    string
	Path     string
}

// MergeHtmlFileCh takes several channels and combine their input
// taken from http://blog.golang.org/pipelines, fan-in, fan-out
func MergeHtmlFileCh(cs ...<-chan HtmlFile) <-chan HtmlFile {
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

// init starts the pool of pipe runners which is the worker pool of
// pandoc processes
func init() {
	piperunner.StartPool()
}
