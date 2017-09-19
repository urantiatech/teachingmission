package db

const SectionsCollection = "sections"

type Section struct {
	Language    string `json:"language"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	ParentId    string `json:"parentid"`
	Description string `json:"description"`
}
