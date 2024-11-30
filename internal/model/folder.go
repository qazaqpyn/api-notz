package model

import "github.com/jackc/pgx/v5/pgtype"

type Folder struct {
	Id        pgtype.UUID        `json:"id,omitempty"`
	Name      string             `json:"name"`
	ParentId  pgtype.UUID        `json:"parent_id"`
	IsRoot    bool               `json:"is_root"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeleteAt  pgtype.Timestamptz `json:"delete_at,omitempty"`
	CreatedBy pgtype.UUID        `json:"created_by"`
}

type UpdateFolder struct {
	Name     string      `json:"name"`
	ParentId pgtype.UUID `json:"parent_id"`
}
