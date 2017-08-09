package teacher

type Teacher struct {
	Id         string   `json:"_id"`
	Slug       string   `json:"slug"`
	Name       string   `json:"name"`
	OtherNames []string `json:"othernames"`
	Profile    string   `json:"profile"`
}
