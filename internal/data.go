package internal

import "time"

type Source struct {
	ID           int64
	URL          string
	Kind         string
	LastUpdateAt *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Channel struct {
	ID          int64
	Title       string
	Description string
	Link        string
	Items       []Record
}

type Record struct {
	ID          int64
	Title       string
	Link        string
	GUID        string
	Description string
	PubDate     time.Time
}
