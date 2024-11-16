package service

import (
	"context"
	"errors"

	"github.com/qazaqpyn/api-notz/internal/model"
	"github.com/qazaqpyn/api-notz/pkg/repository"
)

type NoteService struct {
	repo *repository.Repository
}

func NewNoteService(repo *repository.Repository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) GetAllNotes(ctx context.Context) ([]*model.Note, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		return nil, errors.New("user not found")
	}

	return s.repo.GetAllNotes(ctx, userId)
}

func (s *NoteService) GetNoteById(ctx context.Context, noteId string) (*model.Note, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		return &model.Note{}, errors.New("user not found")
	}

	return s.repo.GetNoteById(ctx, userId, noteId)
}

func (s *NoteService) CreateNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		return &model.Note{}, errors.New("user not found")
	}

	return s.repo.CreateNote(ctx, userId, note)
}

func (s *NoteService) UpdateNote(ctx context.Context, noteId string, input *model.UpdateNoteInput) error {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		return errors.New("user not found")
	}

	return s.repo.UpdateNote(ctx, userId, noteId, input)
}

func (s *NoteService) DeleteNote(ctx context.Context, noteId string) error {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		return errors.New("user not found")
	}

	return s.repo.DeleteNote(ctx, userId, noteId)
}
