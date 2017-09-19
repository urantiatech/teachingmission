package db

import (
	"time"
)

const Collection = "transcripts"

type Teacher struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

type Person struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

type Group struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

type TranscriptSection struct {
	Toc        string   `bson:"toc"`
	Heading    string   `bson:"heading"`
	Paragraphs []string `bson:"paragraphs"`
}

type TranscriptBody struct {
	Html       string              `bson:"html"`
	Paragraphs []string            `bson:"paragraphs"`
	Sections   []TranscriptSection `bson:"sections"`
}

type Transcript struct {
	Id           string         `bson:"id"`
	Slug         string         `bson:"slug"`
	TrSessionId  string         `bson:"trsessionid"`
	Title        string         `bson:"title"`
	Visibility   bool           `bson:"visibility"`
	Status       string         `bson:"status"`
	Language     string         `bson:"language"`
	Date         time.Time      `bson:"date"`
	Location     string         `bson:"location"`
	Teachers     []Teacher      `bson:"teachers"`
	Receiver     Person         `bson:"receiver"`
	Translator   Person         `bson:"translator"`
	Attendees    []Person       `bson:"attendees"`
	Group        Group          `bson:"group"`
	Categories   []Category     `bson:"categories"`
	ImportDomain string         `bson:"importdomain"`
	ImportUrl    string         `bson:"importurl"`
	ImportDate   time.Time      `bson:"importdate"`
	LastUpdate   time.Time      `bson:"lastupdate"`
	Body         TranscriptBody `bson:"body"`
}

type SearchRecord struct {
	Id         string     `bson:"id"`
	Visibility bool       `bson:"visibility"`
	Status     string     `bson:"status"`
	Title      string     `bson:"title"`
	Date       time.Time  `bson:"date"`
	Teachers   []Teacher  `bson:"teachers"`
	Receiver   Person     `bson:"receiver"`
	Group      Group      `bson:"group"`
	Categories []Category `bson:"categories"`
}
