package reader

import (
	"encoding/xml"
)

type Channel struct {
	XMLName     xml.Name `xml:"rss"`
	Title       string   `xml:"channel>title"`
	Description string   `xml:"channel>description"`
	Link        string   `xml:"channel>link"`
	Items       []Item   `xml:"channel>item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	GUID        string `xml:"guid"`
	PubDate     Date   `xml:"pubDate"`
}
