# rss

A very simple library to parse RSS feeds to Go structs. 

## Features

- Parse all required elements of RSS 2.0 feeds

### Todo

- Parse all optional elements of RSS 2.0 feeds
- Parse all required elements of RSS 1.0 feeds
- Parse all optional elements of RSS 1.0 feeds
- Parse all required elements of Atom feeds
- Parse all optional elements of Atom feeds
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

