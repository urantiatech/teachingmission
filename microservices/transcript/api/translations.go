package api

type TranslationsRequest struct {
	Id string `json:"id"`
}

type TranslationsHeaderFields struct {
	Language   string `json:"language"`
	Title      string `json:"title"`
	Visibility bool   `json:"visibility"`
	Status     string `json:"status"`
}

type TranslationsResponse struct {
	Id     string                     `json:"id"`
	Fields []TranslationsHeaderFields `json:"fields"`
	Err    string                     `json:"err,omitempty"`
}
