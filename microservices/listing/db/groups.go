package db

const GroupsCollection = "groups"

type Group struct {
	Language    string `json:"language"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
