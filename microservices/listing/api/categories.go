package api

type CategoriesRequest struct {
	Language string `json:"language"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parentid"`
}

type Category struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ParentId    string `json:"parentid"`
	Description string `json:"description"`
}

type CategoriesResponse struct {
	Count      int        `json:"count"`
	Categories []Category `json:"categories"`
	Err        string     `json:"err,omitempty"`
}
