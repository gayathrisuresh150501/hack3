package models

type Note struct {
	ID       string `json:"id" bson:"_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	OwnerID  string `json:"owner_id"`
	TeamID   string `json:"team_id"`
	IsShared bool   `json:"is_shared"`
}
