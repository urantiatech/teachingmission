package api

type GroupsRequest struct {
	Language string `json:"language"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

type Group struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GroupsResponse struct {
	Count  int     `json:"count"`
	Groups []Group `json:"groups"`
	Err    string  `json:"err,omitempty"`
}
