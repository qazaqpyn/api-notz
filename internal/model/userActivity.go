package model

type UserActivity struct {
	Id        string `json:"id,omitempty"`
	Activity  string `json:"activity"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
