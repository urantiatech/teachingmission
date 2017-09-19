package api

type TaptargetRequest struct {
	Language  string `json:"language"`
	Text      string `json:"text"`
	Tag       string `json:"tag"`
	Css       string `json:"css"`
	Duplicate bool   `json:"duplicate"`
	Value     bool   `json:"value"`
}

type TaptargetResponse struct {
	Text       string      `json:"text"`
	Count      int         `json:"count"`
	Dictionary *Dictionary `json:"dictionary"`
	Err        string      `json:"err,omitempty"`
}
