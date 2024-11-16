package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type RefreshSession struct {
	Id        pgtype.UUID
	UserId    pgtype.UUID
	Token     string
	ExpiresAt time.Time
}
