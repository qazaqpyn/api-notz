package repository

import (
	"context"
	"database/sql"

	"github.com/qazaqpyn/api-notz/internal/model"
)

type TokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{db}
}

func (r *TokenRepository) Create(ctx context.Context, token model.RefreshSession) error {
	return nil
}

func (r *TokenRepository) Get(ctx context.Context, token string) (model.RefreshSession, error) {
	return model.RefreshSession{}, nil
}
