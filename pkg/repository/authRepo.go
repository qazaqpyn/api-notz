package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/qazaqpyn/api-notz/internal/model"
)

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) CreateUser(ctx context.Context, user model.User) error {
	return nil
}

func (r *authRepository) GetUser(ctx context.Context, username, password string) (model.User, error) {
	return model.User{}, nil
}
