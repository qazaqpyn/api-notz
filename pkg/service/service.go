package service

import (
	"context"

	"github.com/qazaqpyn/api-notz/internal/model"
	"github.com/qazaqpyn/api-notz/pkg/repository"
)

type Authorization interface {
	CreateUser(ctx context.Context, user model.User) error
	Login(ctx context.Context, body model.LoginRequest) (string, string, error)
	ParseToken(ctx context.Context, token string) (string, error)
	RefreshTokens(ctx context.Context, token string) (string, string, error)
}

type Note interface {
	GetAllNotes(ctx context.Context) ([]*model.Note, error)
	GetNoteById(ctx context.Context, noteId string) (*model.Note, error)
	CreateNote(ctx context.Context, note *model.Note) (*model.Note, error)
	UpdateNote(ctx context.Context, noteId string, input *model.UpdateNoteInput) error
	DeleteNote(ctx context.Context, noteId string) error
}

type Service struct {
	Authorization
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Note:          NewNoteService(repos),
	}
}
