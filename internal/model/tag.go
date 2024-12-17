package model

import (
	"errors"
	"strings"

	"github.com/lib/pq"
)

type Tag struct {
	Id        string      `json:"id,omitempty" db:"id"`
	Name      string      `json:"name" db:"name"`
	CreatedAt string      `json:"createdAt" db:"created_at"`
	UpdatedAt string      `json:"updatedAt" db:"updated_at"`
	DeleteAt  pq.NullTime `json:"deleteAt,omitempty" db:"deleted_at"`
	CreatedBy string      `json:"createdBy" db:"created_by"`
}

type TagInput struct {
	Name   string `json:"name" db:"name"`
	UserId string `json:"userId,omitempty" db:"user_id"`
	Id     string `json:"id,omitempty" db:"id"`
}

type TagsInput struct {
	Names []string `json:"names"`
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
