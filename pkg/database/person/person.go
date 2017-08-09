package person

type Person struct {
	Id      string `json:"_id"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}
