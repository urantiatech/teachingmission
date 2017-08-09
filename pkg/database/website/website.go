package website

type Website struct {
	Id        string   `json:"id"`
	Domain    string   `json:"domain"`
	Mail      string   `json:"mail"`
	Languages []string `json:"languages"`
}
