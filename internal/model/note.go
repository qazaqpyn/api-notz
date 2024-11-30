package model

import "github.com/jackc/pgx/v5/pgtype"

type Note struct {
	Id         pgtype.UUID        `json:"id,omitempty"`
	Title      string             `json:"title"`
	Body       string             `json:"body"`
	Summary    string             `json:"summary"`
	Transcript string             `json:"transcript"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
	DeleteAt   pgtype.Timestamptz `json:"delete_at,omitempty"`
	CreatedBy  pgtype.UUID        `json:"created_by"`
	UpdatedBy  pgtype.UUID        `json:"updated_by"`
	DeletedBy  pgtype.UUID        `json:"deleted_by,omitempty"`
}

type UpdateNoteInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateNoteTagsInput struct {
	Added   []pgtype.UUID
	Deleted []pgtype.UUID
}
