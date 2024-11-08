package models

import "time"

type Note struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	Owner_ID   string    `json:"owner" bson:"owner"`
	Title      string    `json:"title" bson:"title"`
	Content    string    `json:"content" bson:"content"`
	SharedWith []Share   `json:"shared_with" bson:"shared_with"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

type Share struct {
	UserID     string `json:"user_id" bson:"user_id"`
	Permission string `json:"permission" bson:"permission"` // "view" or "edit"
}

type Plan struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	PlanType string `json:"plan_type" bson:"plan_type"`
	Owner_ID string `json:"owner" bson:"owner"`
}
