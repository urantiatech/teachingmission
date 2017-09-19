package api

type ArticleRequest struct {
	Language string `json:"language"`
	Id       string `json:"id"`
}

type ArticleResponse struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Err   string `json:"err,omitempty"`
}
