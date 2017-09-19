package db

const BeingsCollection = "beings"

type Being struct {
	Language    string `json:"language"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
