package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type RefreshSession struct {
	Id        pgtype.UUID `json:"id,omitempty"`
	UserId    pgtype.UUID `json:"user_id"`
	Token     string      `json:"token"`
	ExpiresAt time.Time   `json:"expires_at"`
}
