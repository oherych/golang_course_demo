package internal

import "time"

type Channel struct {
	Title       string
	Description string
	Link        string
	Items       []Record
}

type Record struct {
	Title       string
	Link        string
	GUID        string
	Description string
	PubDate     time.Time
}
