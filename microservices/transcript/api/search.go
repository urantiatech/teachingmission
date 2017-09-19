package api

import (
	"time"
)

type SearchRequest struct {
	Language   string `json:"language"`
	Visibility string `json:"visibility"`
	Status     string `json:"status"`
	Query      string `json:"query"`
	Category   string `json:"category"`
	Group      string `json:"group"`
	Teacher    string `json:"teacher"`
	Being      string `json:"being"`
	Receiver   string `json:"receiver"`
	StartDate  string `json:"startdate"`
	EndDate    string `json:"enddate"`
	Limit      int    `json:"limit"`
	Skip       int    `json:"skip"`
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Teacher struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Person struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SearchRecord struct {
	Id         string     `json:"id"`
	Visibility bool       `json:"visibility"`
	Status     string     `json:"status"`
	Title      string     `json:"title"`
	Date       time.Time  `json:"date"`
	Teachers   []Teacher  `json:"teachers"`
	Receiver   Person     `json:"receiver"`
	Group      Group      `json:"group"`
	Categories []Category `json:"categories"`
}

type SearchResults struct {
	Results []SearchRecord `json:"results"`
	Limit   int            `json:"limit"`
	Skip    int            `json:"skip"`
	Total   int            `json:"total"`
	Err     string         `json:"err,omitempty"`
}
