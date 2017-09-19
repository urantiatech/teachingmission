package api

type SectionsRequest struct {
	Language string `json:"language"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parentid"`
}

type Section struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ParentId    string `json:"parentid"`
	Description string `json:"description"`
}

type SectionsResponse struct {
	Count    int       `json:"count"`
	Sections []Section `json:"sections"`
	Err      string    `json:"err,omitempty"`
}
