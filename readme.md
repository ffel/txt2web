# Text to Web

The objective of this project is to create a web site based upon
a directory of `txt` files.

This is work in progress!

It uses [pandoc](http://pandoc.org/) to convert [markdown](http://daringfireball.net/projects/markdown/) to html.

The current state is that the following tree

```tree
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

```tree
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