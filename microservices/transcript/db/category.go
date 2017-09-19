package db

type Category struct {
	Id   string `json:"_id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}
