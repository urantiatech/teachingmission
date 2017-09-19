package api

type DefinitionRequest struct {
	Language string `json:"language"`
	Term     string `json:"term"`
}

type DefinitionResponse struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Err        string `json:"err,omitempty"`
}
