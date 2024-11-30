package model

import "github.com/jackc/pgx/v5/pgtype"

type Tag struct {
	Id        pgtype.UUID        `json:"id,omitempty"`
	Name      string             `json:"name"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeleteAt  pgtype.Timestamptz `json:"delete_at,omitempty"`
	CreatedBy pgtype.UUID        `json:"created_by"`
}

type UpdateTagInput struct {
	Name string `json:"name"`
}
