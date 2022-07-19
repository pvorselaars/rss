package rss

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Feed struct {
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	// required elements
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`

	// almost all optional elements; `docs`, `cloud` and `textinput` ignored
	Language  string   `xml:"language"`
	Copyright string   `xml:"copyright"`
	Editor    string   `xml:"managingEditor"`
	Webmaster string   `xml:"webMaster"`
	Date      string   `xml:"pubDate"`
	BuildDate string   `xml:"lastBuildDate"`
	Category  string   `xml:"category"`
	Generator string   `xml:"generator"`
	TTL       int      `xml:"ttl"`
	SkipHours []int    `xml:"skipHours>hour"`
	SkipDays  []string `xml:"skipDays>day"`
	Image     Image    `xml:"image"`
}

type Image struct {
	URL         string `xml:"url"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Width       uint   `xml:"width"`
	Height      uint   `xml:"height"`
}

type Item struct {
	Title       string      `xml:"title"`
	Link        string      `xml:"link"`
	Description string      `xml:"description"`
	Date        string      `xml:"pubDate"`
	Author      string      `xml:"author"`
	Category    []string    `xml:"category"`
	Comments    string      `xml:"comments"`
	Enclosures  []Enclosure `xml:"enclosure"`
	Guid        string      `xml:"guid"`
	Source      Source      `xml:"source"`
}

type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length uint   `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Source struct {
	URL  string `xml:"url,attr"`
	Name string `xml:",innerxml"`
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
	if !strings.Contains(string(data), "<rss") {
		return nil, fmt.Errorf("No RSS 2.0 feed found.")
	}

	var f Feed
	if err := xml.Unmarshal(data, &f); err != nil {
		panic(err)
	}

	if f.Channel == nil {
		return nil, fmt.Errorf("No channel found.")
	}

	return f.Channel, nil
}
