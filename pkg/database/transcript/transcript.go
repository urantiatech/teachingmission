package transcript

import (
	"time"

	_ "github.com/urantiatech/teachingmission/pkg/database"
	"github.com/urantiatech/teachingmission/pkg/database/category"
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

const Collection = "transcripts"

type Teacher2 struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

type Person2 struct {
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
	Id           string              `bson:"id"`
	Slug         string              `bson:"slug"`
	TrSessionId  string              `bson:"trsessionid"`
	Title        string              `bson:"title"`
	Visibility   bool                `bson:"visibility"`
	Status       string              `bson:"status"`
	Language     string              `bson:"language"`
	Date         time.Time           `bson:"date"`
	Location     string              `bson:"location"`
	Teachers     []Teacher2          `bson:"teachers"`
	Receiver     Person2             `bson:"receiver"`
	Translator   Person2             `bson:"translator"`
	Attendees    []Person2           `bson:"attendees"`
	Group        string              `bson:"group"`
	Categories   []category.Category `bson:"categories"`
	ImportDomain string              `bson:"importdomain"`
	ImportUrl    string              `bson:"importurl"`
	ImportDate   time.Time           `bson:"importdate"`
	LastUpdate   time.Time           `bson:"lastupdate"`
	Body         TranscriptBody      `bson:"body"`
}

type SearchRecord struct {
	Id         string              `bson:"id"`
	Language   string              `bson:"language"`
	Visibility bool                `bson:"visibility"`
	Status     string              `bson:"status"`
	Title      string              `bson:"title"`
	Date       time.Time           `bson:"date"`
	Teachers   []Teacher2          `bson:"teachers"`
	Receiver   Person2             `bson:"receiver"`
	Group      string              `bson:"group"`
	Categories []category.Category `bson:"categories"`
}
