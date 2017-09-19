package api

type SearchRequest struct {
	Language   string `json:"language"`
	Visibility string `json:"visibility"`
	Status     string `json:"status"`
	Query      string `json:"query"`
	Category   string `json:"category"`
	Group      string `json:"group"`
	Teacher    string `json:"teacher"`
	Being      string `json:"being"`
	Receiver   string `json:"receiver"`
	StartDate  string `json:"startdate"`
	EndDate    string `json:"enddate"`
	Limit      int    `json:"limit"`
	Skip       int    `json:"skip"`
}
