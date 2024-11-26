package model

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	Id        pgtype.UUID        `json:"id,omitempty"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Email     string             `json:"email" binding:"required"`
	Password  string             `json:"password"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeleteAt  pgtype.Timestamptz `json:"delete_at,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required, gte=6"`
}
