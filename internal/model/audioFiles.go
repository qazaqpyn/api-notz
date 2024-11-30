package model

import "github.com/jackc/pgx/v5/pgtype"

type AudioFiles struct {
	Id        pgtype.UUID        `json:"id,omitempty"`
	FileName  string             `json:"file_name"`
	FilePath  string             `json:"file_path"`
	NoteId    pgtype.UUID        `json:"note_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeleteAt  pgtype.Timestamptz `json:"delete_at,omitempty"`
}
