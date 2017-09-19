package api

type ListingsRequest struct {
	Language   string `json:"language"`
	Categories bool   `json:"categories"`
	Group      bool   `json:"groups"`
	Teachers   bool   `json:"teachers"`
	Beings     bool   `json:"beings"`
	Receivers  bool   `json:"receivers"`
	Sections   bool   `json:"sections"`
}

type ListingsResponse struct {
	Categories CategoriesResponse `json:"categories"`
	Groups     GroupsResponse     `json:"groups"`
	Teachers   TeachersResponse   `json:"teachers"`
	Beings     BeingsResponse     `json:"beings"`
	Receivers  ReceiversResponse  `json:"receivers"`
	Sections   SectionsResponse   `json:"sections"`
	Err        string             `json:"err,omitempty"`
}
