package definition

type Definition struct {
	Id         string `json:"_id"`
	Slug       string `json:"slug"`
	Term       string `json:"term"`
	Definition string `json:"definition"`
}
