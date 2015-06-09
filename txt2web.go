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
	// don't use Section in the key for this makes the key change
	// when the order in the original text file changes, and that's too fragile
	return fmt.Sprintf("#%s", filepath.Join(c.Path, c.PandocId))
}

// String produces the markdown link for Chunk
func (c Chunk) String() string {
	return fmt.Sprintf("[%s](%s)", c.Title, c.Webkey())
}

// txtRoot is relative or absolute path to root of txt2web project
var txtRoot string

// HtmlFile is the contents and the file to write
type HtmlFile struct {
	Contents []byte
	Path     string
}

// Convert does the complete txt2web conversion
func Convert(path string) <-chan HtmlFile {
	// this creates the entire pipeline
	return WriteHtml(Generate(TxtFiles(path)))
}

////////////////////////////////////////////////////////////////////////////////

// Header describes a <h1> chunk of html
type Header struct {
	Header      string
	HeaderLevel int
	Key         string
	Path        string
}

// String prints Header
func (h Header) String() string {
	// you don't want the whitespace here
	return fmt.Sprintf("[%s](%s)", h.Header, h.WebKey())
}

// WebKey creates the inter txt file web key
func (h Header) WebKey() string {
	return fmt.Sprintf("#%s/%s", h.Path, h.Key)
}

/*
// Walk start recursive iteration over sub dir tree
func Walk(root, path string) {
	for _, h := range Headers(root, path) {
		// printing c will be added later, or better, accept a function
		_, err := Contents(h)

		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s+ %v\n", strings.Repeat("  ", h.HeaderLevel-1), h)
		}
	}

	for _, s := range SubDirs(path) {
		Walk(root, s)
	}
}

// Headers gets all headers in path
func Headers(root, path string) []Header {
	files, err := ioutil.ReadDir(path)

	result := make([]Header, 0)

	if err != nil {
		log.Println(err)
		return result
	}

	for _, f := range files {
		if f.Mode().IsRegular() && filepath.Ext(f.Name()) == ".txt" {
			result = append(result, collectheaders(root, filepath.Join(path, f.Name()))...)
		}
	}

	return result
}

// Contents gets the contents that goes with header
func Contents(header Header) (string, error) {
	return "", nil
}

// SubDir gets all sub directories in path such that these can be used in Headers()
func SubDirs(path string) []string {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return make([]string, 0)
	}

	result := make([]string, 0, len(files))

	for _, f := range files {
		if f.IsDir() {
			result = append(result, filepath.Join(path, f.Name()))
		}
	}

	return result
}
*/
