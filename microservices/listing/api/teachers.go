package api

type TeachersRequest struct {
	Language string `json:"language"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	BeingId  string `json:"beingid"`
}

type Teacher struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	OtherNames []string `json:"othernames"`
	BeingId    string   `json:"beingid"`
	Profile    string   `json:"profile"`
}

type TeachersResponse struct {
	Count    int       `json:"count"`
	Teachers []Teacher `json:"teachers"`
	Err      string    `json:"err,omitempty"`
}
