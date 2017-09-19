package db

const ReceiversCollection = "receivers"

type Receiver struct {
	Language string `json:"language"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Profile  string `json:"profile"`
}
