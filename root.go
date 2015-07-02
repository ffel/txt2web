package txt2web

import (
	"bytes"
	"log"
	"path/filepath"
	"text/template"
)

type AngRoutes struct {
	Key string
}

// The node `WriteRoot` adds the files in the root of the web site.

func WriteRoot(in <-chan Chunk) <-chan HtmlFile {
	out := make(chan HtmlFile)

	go func() {
		// no need to wait for messages on in, we can write index.html right away
		out <- HtmlFile{
			Contents: []byte(index),
			Title:    "index",
			Path:     filepath.Join("..", "index.html"),
		}

		// same for styles
		out <- HtmlFile{
			Contents: []byte(css),
			Title:    "styles",
			Path:     filepath.Join("..", "pandoc.css"),
		}

		var keys []AngRoutes

		for c := range in {
			keys = append(keys, AngRoutes{Key: c.Webkey()})
		}

		t := template.New("routes")
		t, err := t.Parse(appjs)
		if err != nil {
			log.Fatal(err)
		}

		buff := bytes.NewBufferString("")

		err = t.Execute(buff, struct{ Routes []AngRoutes }{keys})

		if err != nil {
			log.Println("template:", err)
		}

		out <- HtmlFile{
			Contents: buff.Bytes(),
			Title:    "angular app",
			Path:     filepath.Join("..", "app.js"),
		}

		close(out)
	}()

	return out
}

const index = `<!DOCTYPE html>
<html ng-app="myApp">
<head>
  <meta charset="utf-8">
  <meta name="generator" content="pandoc">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=yes">
  <title></title>
  <style type="text/css">code{white-space: pre;}</style>
  <!--[if lt IE 9]>
    <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
  <![endif]-->
  <style type="text/css">
table.sourceCode, tr.sourceCode, td.lineNumbers, td.sourceCode {
  margin: 0; padding: 0; vertical-align: baseline; border: none; }
table.sourceCode { width: 100%; line-height: 100%; }
td.lineNumbers { text-align: right; padding-right: 4px; padding-left: 4px; color: #aaaaaa; border-right: 1px solid #aaaaaa; }
td.sourceCode { padding-left: 5px; }
code > span.kw { color: #007020; font-weight: bold; }
code > span.dt { color: #902000; }
code > span.dv { color: #40a070; }
code > span.bn { color: #40a070; }
code > span.fl { color: #40a070; }
code > span.ch { color: #4070a0; }
code > span.st { color: #4070a0; }
code > span.co { color: #60a0b0; font-style: italic; }
code > span.ot { color: #007020; }
code > span.al { color: #ff0000; font-weight: bold; }
code > span.fu { color: #06287e; }
code > span.er { color: #ff0000; font-weight: bold; }
  </style>
  <link rel="stylesheet" href="pandoc.css">
    <script src="//code.angularjs.org/1.3.0-rc.1/angular.min.js"></script>
    <script src="//code.angularjs.org/1.3.0-rc.1/angular-route.min.js"></script>
    <script src="app.js"></script>
</head>
<body>

    
    <div class="container">
        <div ng-view></div>
    </div>

</body>
</html>
`

const css = `body {
    margin: auto;
    padding-right: 1em;
    padding-left: 1em;
    max-width: 44em; 
    border-left: 1px solid black;
    border-right: 1px solid black;
    color: black;
    font-family: Verdana, sans-serif;
    font-size: 100%;
    line-height: 140%;
    color: #333; 
}
pre {
    border: 1px dotted gray;
    background-color: #ececec;
    color: #1111111;
    padding: 0.5em;
}
code {
    font-family: monospace;
}
h1 a, h2 a, h3 a, h4 a, h5 a { 
    text-decoration: none;
    color: #7a5ada; 
}
h1, h2, h3, h4, h5 { font-family: verdana;
                     font-weight: bold;
                     border-bottom: 1px dotted black;
                     color: #7a5ada; }
h1 {
        font-size: 130%;
}

h2 {
        font-size: 110%;
}

h3 {
        font-size: 95%;
}

h4 {
        font-size: 90%;
        font-style: italic;
}

h5 {
        font-size: 90%;
        font-style: italic;
}

h1.title {
        font-size: 200%;
        font-weight: bold;
        padding-top: 0.2em;
        padding-bottom: 0.2em;
        text-align: left;
        border: none;
}

dt code {
        font-weight: bold;
}
dd p {
        margin-top: 0;
}

#footer {
        padding-top: 1em;
        font-size: 70%;
        color: gray;
        text-align: center;
        }
`

const appjs = `
var myApp = angular.module('myApp', ['ngRoute']);

myApp.config(function ($routeProvider) {
    
    $routeProvider

    .when('/', {
        templateUrl: 'pages/index.html',
        controller: 'mctrl'
    })
    
    {{with .Routes}}
    {{range .}}
	.when('/{{.Key}}', {
	    templateUrl: 'pages/{{.Key}}.html',
	    controller: 'mctrl'
	})
    {{end}}
    {{end}}
   
});

myApp.controller('mctrl', ['$scope', '$log', function($scope, $log) {
    
    $scope.name = 'Main';
    
}]);
`
