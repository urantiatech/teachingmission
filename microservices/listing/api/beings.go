package api

type BeingsRequest struct {
	Language string `json:"language"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

type Being struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BeingsResponse struct {
	Count  int     `json:"count"`
	Beings []Being `json:"beings"`
	Err    string  `json:"err,omitempty"`
}
