package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/qazaqpyn/api-notz/internal/model"
)

type TokenRepository struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{db}
}

func (r *TokenRepository) Create(ctx context.Context, token model.RefreshSession) error {
	// Create if it doesn't exist, Update if it exists
	return nil
}

func (r *TokenRepository) Get(ctx context.Context, token string) (model.RefreshSession, error) {
	return model.RefreshSession{}, nil
}
