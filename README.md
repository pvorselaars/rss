# rss

A very simple library to parse RSS feeds to Go structs. 

## Features

- Parse [RSS 2.0](https://www.rssboard.org/rss-specification) feeds

### Todo

- Parse [RSS 1.0](https://web.resource.org/rss/1.0/spec) feeds
- Parse [Atom](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc4287.html) feeds
- Automated tests

## Usage

Using this library is extremely simple:

```
package main

import "github.com/pvorselaars/rss"

func main(){

  channel, err := rss.Fetch("https://example.com/rss")
  if err != nil {
    // handle error
  }

  // use channel struct
}

```

This library uses the following structs for channels and items.

```
type Channel struct {
	Title       string
	Link        string
	Description string
	Items       []Item
}

type Item struct {
	Title       string
	Link        string
	Description string
	Date        string
	Author      string
	Category    []string
}

```

