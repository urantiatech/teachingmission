package api

type BannerRequest struct {
	Language string `json:"language"`
}

type BannerResponse struct {
	Url  string `json:"url"`
	Text string `json:"text"`
	Err  string `json:"err,omitempty"`
}
