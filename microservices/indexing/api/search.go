package api

type SearchRequest struct {
	Language string `json:"language"`
	Bucket   string `json:"bucket"`
	Query    string `json:"query"`
}

type SearchRecord struct {
	Id     string `json:"id"`
	Bucket string `json:"bucket"`
}

type SearchResponse struct {
	Results []SearchRecord `json:"results"`
	Limit   int            `json:"limit"`
	Skip    int            `json:"skip"`
	Total   int            `json:"total"`
	Err     string         `json:"err,omitempty"`
}
