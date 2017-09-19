package db

const CategoriesCollection = "categories"

type Category struct {
	Language    string `json:"language"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	ParentId    string `json:"parentid"`
	Description string `json:"description"`
}
