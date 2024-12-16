package model

import (
	"errors"

	"github.com/lib/pq"
	"github.com/qazaqpyn/api-notz/internal/tools"
)

type User struct {
	Id        string      `json:"id,omitempty" db:"id"`
	FirstName string      `json:"firstName" db:"first_name"`
	LastName  string      `json:"lastName" db:"last_name"`
	Email     string      `json:"email" binding:"required" db:"email"`
	Password  string      `json:"password" db:"password"`
	CreatedAt string      `json:"createdAt" db:"created_at"`
	UpdatedAt string      `json:"updatedAt" db:"updated_at"`
	DeleteAt  pq.NullTime `json:"deleteAt,omitempty" db:"deleted_at"`
}

type RegisterRequest struct {
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" binding:"required" db:"email"`
	Password  string `json:"password" db:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password"`
}

func (r *LoginRequest) Validate() error {
	if r.Email == "" || r.Password == "" {
		return errors.New("please input email and password")
	}

	return nil
}

func (r *RegisterRequest) Validate() error {
	if r.FirstName == "" {
		return errors.New("please input first name")
	}

	if r.LastName == "" {
		return errors.New("please input last name")
	}

	if !tools.IsEmailValid(r.Email) {
		return errors.New("please input email")
	}

	if r.Password == "" {
		return errors.New("please input password")
	}

	if len(r.Password) < 5 {
		return errors.New("pasword must have at least 5 characters")
	}

	return nil
}
