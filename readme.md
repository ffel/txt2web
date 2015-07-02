Text to Web
===========

Purpose
-------

The objective of this project is to create a web site based upon a
directory of `txt` files.

This is work in progress! (Have a look at [hugo](http://gohugo.io/) if
you're looking for a static site generator.)

It uses [pandoc](http://pandoc.org/) to convert
[markdown](http://daringfireball.net/projects/markdown/) to html.

The current state is that the following tree

``` {.tree}
example
├── dira
│   ├── filec.txt
│   └── filed.txt
├── dirb
│   └── filee.txt
├── filea.txt
└── fileb.txt
```

is transformed into the following html tree

``` {.tree}
example_html/
└── pages
    ├── dira
    │   ├── vier-nulla-euismod-placerat-nunc-at-mattis.html
    │   ├── vijf-donec-lacus-leo.html
    │   ├── zes-fusce-non-aliquet-tortor..html
    │   └── zeven-nulla-ut-faucibus-felis.html
    ├── dirb
    │   ├── acht-pellentesque-lacinia.html
    │   ├── negen-vivamus-eget-cursus-erat-in-pharetra-neque.html
    │   └── tien-phasellus-lorem-eros.html
    ├── drie-pellentesque-lobortis-lacus.html
    ├── een-lorem-ipsum-dolor-sit-amet.html
    └── twee-morbi-finibus-rutrum-condimentum..html
```

Pandoc has *no* means to deal with links between local files. You have
to use angular routes `#<path><anchor>` where `<path>` is the path
relative to the web root and where `<anchor>` is the anchor [provided by
pandoc](http://pandoc.org/README.html#internal-links) (typically, the
section title in lower case and spaces replaced for dashes).

However, internal links to first order sections in one `txt` file that
will end up in several `html` files will be dealt with as these will be
replaced for angular routes.

Commands
--------

`t2w` is the main command and takes a tree of `txt` files and writes a
tree of `html` files.

`t2toc` is a utility command that prints a table of contents. The links
can be used in your `txt` files.

> `devserver/server` is a simple file server to test the contents in
> `example_html`

Technical Background
--------------------

This project implements a [pipeline](http://blog.golang.org/pipelines)
approach. The default behaviour is in the `Convert()` function which
chains the following nodes:

1.  `TxtFiles()` walks a directory and finds all `txt` files.

2.  `Generate()` reads the `txt` files and generates objects that
    contain a json representation of the text.

3.  `Index()` generates `index.txt` and therefore `index.html` in each
    directory that displays `txt` sections in the current directory as
    well a list of subdirectories.

4.  `References()` replaces in-file references for references that work
    between html files.

5.  `Split()` splits each such object such that it contains one first
    level markdown section.

6.  `WriteHtml()` converts chunks to html.

7.  `WriteRoot()` adds the web site root contents (`index.html`,
    `app.js`, styles).

Additional nodes not used in the default `t2w`:

1.  `Toc()` collects table of contents information (used in `t2toc`)
