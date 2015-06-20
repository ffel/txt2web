Text to Web
===========

Purpose
-------

The objective of this project is to create a web site based upon a directory of `txt` files.

This is work in progress!

It uses [pandoc](http://pandoc.org/) to convert [markdown](http://daringfireball.net/projects/markdown/) to html.

The current state is that the following tree

``` tree
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

``` tree
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

Technical Background
--------------------

This project implements a [pipeline](http://blog.golang.org/pipelines) approach. The default behaviour is in the `Convert()` function which chains the following nodes:

1.  `TxtFiles()` walks a directory and finds all `txt` files.

2.  `Generate()` reads the `txt` files and generates objects that contain a json representation of the text.

3.  `Split()` splits each such object such that it contains one first level markdown section.

4.  `WriteHtml()` converts objects to html.

> github flavoured markdown is obtained with `pandoc readme.md -t markdown_github -o readme.md`.
