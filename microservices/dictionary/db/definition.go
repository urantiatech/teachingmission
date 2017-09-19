package db

const Collection = "definitions"

type Definition struct {
	Language   string `json:"language"`
	Term       string `json:"term"`
	Definition string `json:"definition"`
}
