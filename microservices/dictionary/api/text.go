package api

type TextRequest struct {
	Language  string `json:"language"`
	Text      string `json:"text"`
	Tag       string `json:"tag"`
	Css       string `json:"css"`
	Duplicate bool   `json:"duplicate"`
	Value     bool   `json:"value"`
}

type TextResponse struct {
	Text  string `json:"text"`
	Count int    `json:"count"`
	Err   string `json:"err,omitempty"`
}
