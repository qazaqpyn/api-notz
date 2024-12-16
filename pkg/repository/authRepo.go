package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/qazaqpyn/api-notz/internal/model"
	"github.com/sirupsen/logrus"
)

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) CreateUser(ctx context.Context, user model.RegisterRequest) error {
	_, err := r.db.NamedExecContext(ctx, `
		INSERT INTO users (
			first_name
			, last_name
			, email
			, password
		)
		VALUES (
			:first_name
			, :last_name
			, :email
			, :password
		)
	`, user)

	return err
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}

	nstmt, err := r.db.PrepareNamedContext(ctx, `
		SELECT * from users
		WHERE email = :email
	`)
	if err != nil {
		return &user, err
	}

	err = nstmt.Get(&user, map[string]interface{}{
		"email": email,
	})
	if err != nil {
		logrus.Errorf(err.Error())
		return &user, errors.New("user not found")
	}

	return &user, nil
}
