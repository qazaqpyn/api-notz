package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/qazaqpyn/api-notz/internal/model"
)

type TagRepository struct {
	db *sqlx.DB
}

func NewTagRepository(db *sqlx.DB) *TagRepository {
	return &TagRepository{db}
}

func (t *TagRepository) GetAllTags(ctx context.Context) ([]model.Tag, error) {
	tags := []model.Tag{}
	err := t.db.SelectContext(ctx, tags, "SELECT * FROM tags ORDER BY updated_by")
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *TagRepository) CreateTags(ctx context.Context, tags []model.TagInput) error {
	_, err := t.db.NamedExecContext(ctx, `
		INSERT INTO tags (name, created_by) 
		VALUES (:name, :user_id)
	`, tags)

	return err
}

func (t *TagRepository) GetUserTags(ctx context.Context, userId string) ([]model.Tag, error) {
	tags := []model.Tag{}
	user := map[string]string{
		"user": userId,
	}

	nstmt, err := t.db.PrepareNamedContext(ctx, `
		SELECT * FROM tags
		WHERE created_by = :user_id
		ORDER BY updated_by
	`)
	if err != nil {
		return nil, err
	}

	err = nstmt.Select(tags, user)

	return tags, err
}

func (t *TagRepository) GetByUserTagById(ctx context.Context, userId string, tagId string) (*model.Tag, error) {
	tag := model.Tag{}

	nstmt, err := t.db.PrepareNamedContext(ctx, `
		SELECT * FROM tags
		WHERE created_by = :user_id
			AND id = :id
	`)
	if err != nil {
		return nil, err
	}

	err = nstmt.Select(&tag, map[string]interface{}{
		"UserId": userId,
		"Id":     tagId,
	})

	return &tag, err
}

func (t *TagRepository) DeleteTag(ctx context.Context, tagId string) error {
	_, err := t.db.NamedExecContext(ctx, `
		DELETE FROM tags WHERE id = :id
	`, map[string]interface{}{
		"id": tagId,
	})

	return err
}

func (t *TagRepository) UpdateTag(ctx context.Context, input *model.TagInput) error {
	if input.Id == "" {
		return errors.New("tag id is empty")
	}

	_, err := t.db.NamedExecContext(ctx, `
		UPDATE tags
		SET
			name = :name
		WHERE id = :id
	`, input)

	return err
}
