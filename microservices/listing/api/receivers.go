package api

type ReceiversRequest struct {
	Language string `json:"language"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

type Receiver struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

type ReceiversResponse struct {
	Count     int        `json:"count"`
	Receivers []Receiver `json:"receivers"`
	Err       string     `json:"err,omitempty"`
}
