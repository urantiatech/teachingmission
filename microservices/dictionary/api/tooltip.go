package api

type TooltipRequest struct {
	Language  string `json:"language"`
	Text      string `json:"text"`
	Css       string `json:"css"`
	Duplicate bool   `json:"duplicate"`
}

type TooltipResponse struct {
	Text  string `json:"text"`
	Count int    `json:"count"`
	Err   string `json:"err,omitempty"`
}
