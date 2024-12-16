package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/qazaqpyn/api-notz/internal/model"
)

type Authorization interface {
	CreateUser(ctx context.Context, user model.RegisterRequest) error
	GetUserByEmail(ctx context.Context, username string) (*model.User, error)
}

type Note interface {
	GetAllNotes(ctx context.Context, userId string) ([]*model.Note, error)
	GetNoteById(ctx context.Context, userId, noteId string) (*model.Note, error)
	CreateNote(ctx context.Context, userId string, note *model.Note) (*model.Note, error)
	UpdateNote(ctx context.Context, userId, noteId string, input *model.UpdateNoteInput) error
	DeleteNote(ctx context.Context, userId, noteId string) error
}

type Token interface {
	Create(ctx context.Context, token model.RefreshSession) error
	Get(ctx context.Context, token string) (model.RefreshSession, error)
}

type Tag interface {
	GetAllTags(ctx context.Context) ([]model.Tag, error)
	CreateTags(ctx context.Context, tags []model.TagInput) error
	GetUserTags(ctx context.Context, userId string) ([]model.Tag, error)
	DeleteTag(ctx context.Context, tagId string) error
	UpdateTag(ctx context.Context, input *model.TagInput) error
	GetByUserTagById(ctx context.Context, userId string, tagId string) (*model.Tag, error)
}

type Repository struct {
	Authorization
	Note
	Token
	Tag
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Note:          NewNoteRepository(db),
		Token:         NewTokenRepository(db),
		Tag:           NewTagRepository(db),
	}
}
