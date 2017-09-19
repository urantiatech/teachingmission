package api

import (
	"time"
)

type TranscriptRequest struct {
	Language string `json:"language"`
	Id       string `json:"id"`
}

type TranscriptSection struct {
	Toc        string   `json:"toc"`
	Heading    string   `json:"heading"`
	Paragraphs []string `json:"paragraphs"`
}

type TranscriptBody struct {
	Html       string              `json:"html"`
	Paragraphs []string            `json:"paragraphs"`
	Sections   []TranscriptSection `json:"sections"`
}

type TranscriptResponse struct {
	Id           string         `json:"id"`
	Slug         string         `json:"slug"`
	TrSessionId  string         `json:"trsessionid"`
	Title        string         `json:"title"`
	Visibility   bool           `json:"visibility"`
	Status       string         `json:"status"`
	Language     string         `json:"language"`
	Date         time.Time      `json:"date"`
	Location     string         `json:"location"`
	Teachers     []Teacher      `json:"teachers"`
	Receiver     Person         `json:"receiver"`
	Translator   Person         `json:"translator"`
	Attendees    []Person       `json:"attendees"`
	Group        Group          `json:"group"`
	Categories   []Category     `json:"categories"`
	ImportDomain string         `json:"importdomain"`
	ImportUrl    string         `json:"importurl"`
	ImportDate   time.Time      `json:"importdate"`
	LastUpdate   time.Time      `json:"lastupdate"`
	Body         TranscriptBody `json:"body"`
	Err          string         `json:"err,omitempty"`
}
