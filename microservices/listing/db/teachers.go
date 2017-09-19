package db

const TeachersCollection = "teachers"

type Teacher struct {
	Language   string   `json:"language"`
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	OtherNames []string `json:"othernames"`
	BeingId    string   `json:"beingid"`
	Profile    string   `json:"profile"`
}
