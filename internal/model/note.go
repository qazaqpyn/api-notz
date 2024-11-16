package model

import "github.com/jackc/pgx/v5/pgtype"

type Note struct {
	Id      pgtype.UUID `json:"id", omitempty`
	Title   string      `json:"title"`
	Content string      `json:"content"`
	UserId  pgtype.UUID `json:"user_id"`
}

type UpdateNoteInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
