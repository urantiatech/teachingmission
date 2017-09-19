package api

type ListingsRequest struct {
	Language   string `json:"language"`
	Categories bool   `json:"categories"`
	Groups     bool   `json:"groups"`
	Teachers   bool   `json:"teachers"`
	Beings     bool   `json:"beings"`
	Receivers  bool   `json:"receivers"`
	Sections   bool   `json:"sections"`
}
