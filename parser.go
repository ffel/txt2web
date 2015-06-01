package txt2web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ffel/pandocfilter"
	"github.com/ffel/piperunner"
)

func init() {
	// prepare a pool of pandoc runners in the background
	piperunner.StartPool()
}

func collectheaders(root, file string) []Header {
	result := make([]Header, 0)
	jsondata, err := getJson(file)

	if err != nil {
		return result
	}

	// remove root from file
	rel, err := filepath.Rel(root, file)

	if err != nil {
		log.Println(root, file, err)
		return result
	}

	f := &pdtoc{file: rel, headers: make([]Header, 0)}

	pandocfilter.Walk(f, jsondata)

	return f.headers
}

// function getJson converts contents of file to pandoc json
func getJson(file string) (interface{}, error) {
	txt, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	resultc := piperunner.Exec("pandoc -f markdown -t json", txt)

	result := <-resultc

	if result.Err != nil {
		return nil, result.Err
	}

	var jsondata interface{}
	err = json.Unmarshal(result.Text, &jsondata)

	if err != nil {
		return nil, err
	}

	return jsondata, nil
}

// type pdtoc implements pandoc runner filter and retrieves headers
type pdtoc struct {
	file    string
	headers []Header
}

func (p *pdtoc) Value(level int, key string, value interface{}) (bool, interface{}) {

	ok, t, c := pandocfilter.IsTypeContents(value)

	if ok && t == "Header" {
		title, key, lev := p.processHeader(c)
		p.headers = append(p.headers, Header{Header: title, HeaderLevel: lev, Key: key, Path: p.file})
	}

	return true, value
}

func (p *pdtoc) processHeader(json interface{}) (title, key string, level int) {
	lev, err := pandocfilter.GetNumber(json, "0")

	if err != nil {
		return "", "", 0
	}

	ref, err := pandocfilter.GetString(json, "1", "0")

	if err != nil {
		return "", "", 0
	}

	label, err := pandocfilter.GetObject(json, "2")

	if err != nil {
		return "", "", 0
	}

	col := &collector{}

	pandocfilter.Walk(col, label)

	return col.value, ref, int(lev)
}

// collector walks the header c and collects the Str
type collector struct {
	value string
}

func (coll *collector) Value(level int, key string, value interface{}) (bool, interface{}) {
	ok, t, c := pandocfilter.IsTypeContents(value)

	if ok && t == "Str" {
		coll.value += c.(string) + " "
	}

	return true, value
}
