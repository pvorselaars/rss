package rss

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Feed struct {
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Date        string   `xml:"pubDate"`
	Author      string   `xml:"creator"`
	Category    []string `xml:"category"`
}

func Fetch(url string) (*Channel, error) {
	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ch, err := Parse(body)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func Parse(data []byte) (*Channel, error) {
	var f Feed
	if err := xml.Unmarshal(data, &f); err != nil {
		panic(err)
	}

	if f.Channel == nil {
		return nil, fmt.Errorf("No channel found in %q", string(data))
	}

	return f.Channel, nil
}
