package model

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	Id       pgtype.UUID `json:"id", omitempty`
	Username string      `json:"username"`
	Email    string      `json:"email" binding:"required"`
	Password string      `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required, gte=3"`
	Password string `json:"password" binding:"required, gte=6"`
}
