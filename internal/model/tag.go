package model

import (
	"errors"
	"strings"
)

type Tag struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeleteAt  string `json:"delete_at,omitempty"`
	CreatedBy string `json:"created_by"`
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
