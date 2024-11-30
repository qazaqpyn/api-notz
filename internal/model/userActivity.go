package model

import "github.com/jackc/pgx/v5/pgtype"

type UserActivity struct {
	Id        pgtype.UUID        `json:"id,omitempty"`
	Activity  string             `json:"activity"`
	UserId    pgtype.UUID        `json:"user_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}
