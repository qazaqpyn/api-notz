package repository

import (
	"context"
	"database/sql"

	"github.com/qazaqpyn/api-notz/internal/model"
)

type NoteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) *NoteRepository {
	return &NoteRepository{db}
}

func (r *NoteRepository) GetAllNotes(ctx context.Context, userId string) ([]*model.Note, error) {
	return nil, nil
}

func (r *NoteRepository) GetNoteById(ctx context.Context, userId, noteId string) (*model.Note, error) {
	return &model.Note{}, nil
}

func (r *NoteRepository) CreateNote(ctx context.Context, userId string, note *model.Note) (*model.Note, error) {
	return &model.Note{}, nil
}

func (r *NoteRepository) UpdateNote(ctx context.Context, userId, noteId string, input *model.UpdateNoteInput) error {
	return nil
}

func (r *NoteRepository) DeleteNote(ctx context.Context, userId, noteId string) error {
	return nil
}
