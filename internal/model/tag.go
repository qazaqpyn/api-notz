package model

import (
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

type Tag struct {
	Id        pgtype.UUID        `json:"id,omitempty"`
	Name      string             `json:"name"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	DeleteAt  pgtype.Timestamptz `json:"delete_at,omitempty"`
	CreatedBy pgtype.UUID        `json:"created_by"`
}

type TagInput struct {
	Name   string `json:"name"`
	UserId string `json:"user_id,omitempty"`
	Id     string `json:"id,omitempty"`
}

func (t *TagInput) Validate() error {
	if strings.TrimSpace(t.Name) == "" {
		return errors.New("name is empty")
	}

	return nil
}

func (t *TagInput) AddUserId(userId string) {
	t.UserId = userId
}

func (t *TagInput) AddId(tagId string) {
	t.Id = tagId
}
