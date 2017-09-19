package db

const Collection = "quotes"

type Quote struct {
	Id       string `json:"id"`
	Language string `json:"language"`
	Teacher  string `json:"teacher"`
	Text     string `json:"text"`
}
