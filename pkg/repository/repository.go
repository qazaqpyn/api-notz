package repository

import (
	"context"
	"database/sql"

	"github.com/qazaqpyn/api-notz/internal/model"
)

type Authorization interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, username, password string) (model.User, error)
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
	GetAllTags(ctx context.Context) ([]*model.Tag, error)
	CreateTags(ctx context.Context, tags *model.TagInput) ([]*model.Tag, error)
	GetUserTags(ctx context.Context, userId string) ([]*model.Tag, error)
	DeleteTag(ctx context.Context, tagId string) error
	UpdateTag(ctx context.Context, tagId string, input *model.TagInput) error
}

type Repository struct {
	Authorization
	Note
	Token
	Tag
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Note:          NewNoteRepository(db),
		Token:         NewTokenRepository(db),
		Tag:           NewTagRepository(db),
	}
}
