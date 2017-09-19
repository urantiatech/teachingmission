package api

type TranscriptRequest struct {
	Language string `json:"language"`
	Id       string `json:"id"`
}

type TranslationsRequest struct {
	Id string `json:"id"`
}
