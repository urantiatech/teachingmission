package api

type QuoteRequest struct {
	Language string `json:"language"`
}

type QuoteResponse struct {
	Id       string `json:"id"`
	Language string `json:"language"`
	Teacher  string `json:"teacher"`
	Text     string `json:"text"`
	Err      string `json:"err,omitempty"`
}
