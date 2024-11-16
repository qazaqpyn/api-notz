package repository

import (
	"context"
	"database/sql"

	"github.com/qazaqpyn/api-notz/internal/model"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) CreateUser(ctx context.Context, user model.User) error {
	return nil
}

func (r *authRepository) GetUser(ctx context.Context, username, password string) (model.User, error) {
	return model.User{}, nil
}
