package model

import (
	"time"
)

type RefreshSession struct {
	Id        string    `json:"id,omitempty"`
	UserId    string    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
