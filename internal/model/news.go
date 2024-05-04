package model

import (
	"encoding/xml"
	"time"
)

type Post struct {
	ID      string    `json:"id" db:"id"`
	Title   string    `json:"title" db:"title"`
	Content string    `json:"content" db:"content"`
	PubTime time.Time `json:"pub_time" db:"pub_date"`
	Link    string    `json:"link" db:"link"`
}

type RssPost struct {
	RssVersion xml.Name `xml:"rss"`
	Channel    Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Item        []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
