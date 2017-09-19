package db

const Collection = "articles"

type Article struct {
	Id       string `json:"id"`
	Language string `json:"language"`
	Title    string `json:"title"`
	Text     string `json:"text"`
}
