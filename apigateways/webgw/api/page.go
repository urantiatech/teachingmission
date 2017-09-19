package api

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Menu struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	SubMenu *Menu  `json:"submemnu"`
}

type Banner struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}

type Quote struct {
	Text    string `json:"text"`
	Teacher string `json:"teacher"`
}

type Category struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ParentId    string `json:"parentid"`
	Description string `json:"description"`
}

type Group struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Filter struct {
	Category  string `json:"category"`
	Group     string `json:"group"`
	Teacher   string `json:"teacher"`
	Receiver  string `json:"receiver"`
	Being     string `json:"being"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type Teacher struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	OtherNames []string `json:"othernames"`
	Being      string   `json:"being"`
	Profile    string   `json:"profile"`
}

type Being struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Person struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

type Receiver struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

type TranscriptSection struct {
	TocTitle   string   `json:"toctitle"`
	Heading    string   `json:"heading"`
	Paragraphs []string `json:"paragraphs"`
}

type TranscriptBody struct {
	Html       string              `bson:"html"`
	Paragraphs []string            `bson:"paragraphs"`
	Sections   []TranscriptSection `bson:"sections"`
}

type Transcript struct {
	Id         string         `json:"id"`
	Title      string         `json:"title"`
	Categories []Category     `json:"categories"`
	Date       time.Time      `json:"date"`
	Location   string         `json:"location"`
	Teachers   []Teacher      `json:"teachers"`
	Receiver   Person         `json:"receiver"`
	Attendees  []Person       `json:"attendees"`
	Body       TranscriptBody `json:"body"`
}

type InternalResult struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Teachers []Teacher `json:"teachers"`
	Excerpts string    `json:"excerpts"`
}

type ExternalResult struct {
	Url      string `json:"url"`
	Title    string `json:"title"`
	Excerpts string `json:"excerpts"`
}

type Pager struct {
	Url      string `json:"url"`
	PageNo   int    `json:"pageno"`
	Active   bool   `json:"active"`
	Disabled bool   `json:"disabled"`
}

type Search struct {
	Query           string           `json:"query"`
	Filters         Filter           `json:"filters"`
	Results         []InternalResult `json:"results"`
	ExternalResults []ExternalResult `json:"externalresults"`
	Limit           int              `json:"limit"`
	Skipped         int              `json:"skipped"`
	Total           int              `json:"total"`
}

type Article struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Meta struct {
	Keywords    string `json:"meta"`
	Description string `json:"description"`
}

type TranscriptHeader struct {
	Language   string `json:"language"`
	Title      string `json:"title"`
	Visibility bool   `json:"visibility"`
	Status     string `json:"status"`
}

type Translations struct {
	Id      string             `json:"id"`
	Headers []TranscriptHeader `json:"headers"`
	Err     string             `json:"err"`
}

type CategoriesListing struct {
	Count      int        `json:"count"`
	Categories []Category `json:"categories"`
}

type GroupsListing struct {
	Count  int     `json:"count"`
	Groups []Group `json:"groups"`
}

type TeachersListing struct {
	Count    int       `json:"count"`
	Teachers []Teacher `json:"teachers"`
}

type BeingsListing struct {
	Count  int     `json:"count"`
	Beings []Being `json:"beings"`
}

type ReceiversListing struct {
	Count     int        `json:"count"`
	Receivers []Receiver `json:"receiver"`
}

type Section struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ParentId    string `json:"parentid"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type SectionsListing struct {
	Count    int       `json:"count"`
	Sections []Section `json:"sections"`
}

type Listings struct {
	Categories CategoriesListing `json:"categories"`
	Groups     GroupsListing     `json:"groups"`
	Teachers   TeachersListing   `json:"teachers"`
	Beings     BeingsListing     `json:"beings"`
	Receivers  ReceiversListing  `json:"receivers"`
	Sections   SectionsListing   `json:"archives"`
	Err        string            `json:"err"`
}

type Breadcrumb struct {
	Titles []string `json:"titles"`
}

type Definition struct {
	Term  string `json:"term"`
	Value string `json:"value"`
}

type Dictionary struct {
	Definitions []Definition `json:"definitions"`
}

type Page struct {
	Title        string        `json:"title"`
	Meta         *Meta         `json:"meta"`
	User         *User         `json:"user"`
	Menu         *Menu         `json:"menu"`
	Banner       *Banner       `json:"banner"`
	Quote        *Quote        `json:"quote"`
	Breadcrumb   *Breadcrumb   `json:"breadcrumb"`
	Search       *Search       `json:"search"`
	Translations *Translations `json:"translations"`
	Transcript   *Transcript   `json:"transcript"`
	Article      *Article      `json:"article"`
	Listings     *Listings     `json:"listings"`
	Dictionary   *Dictionary   `json:"dictionary"`
	Err          string        `json:"err"`
}

func ContextVal(ctx context.Context, key string) (string, error) {
	if val := ctx.Value(key); val != nil {
		return fmt.Sprint(val), nil
	}
	return "", errors.New("key not found")
}
