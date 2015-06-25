package txt2web

import (
	"fmt"
	"path/filepath"

	"github.com/ffel/piperunner"
)

// init starts the pool of pipe runners which is the worker pool of
// pandoc processes
func init() {
	piperunner.StartPool()
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

	var htmlc <-chan HtmlFile

	// send file name and contents
	htmlc = WriteHtml(chunkc)

	return htmlc
}
