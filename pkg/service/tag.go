package service

import (
	"context"

	"github.com/qazaqpyn/api-notz/internal/model"
	"github.com/qazaqpyn/api-notz/pkg/repository"
)

type TagService struct {
	repo *repository.Repository
}

func NewTagService(repo *repository.Repository) *TagService {
	return &TagService{repo: repo}
}

func (t *TagService) GetAllTags(ctx context.Context) ([]model.Tag, error) {
	return t.repo.GetAllTags(ctx)
}

func (t *TagService) CreateTags(ctx context.Context, userId string, tags []model.TagInput) error {
	// Add userId to Input
	t.addUserIdToInput(tags, userId)

	return t.repo.CreateTags(ctx, tags)
}

func (t *TagService) GetUserTags(ctx context.Context, userId string) ([]model.Tag, error) {
	return t.repo.GetUserTags(ctx, userId)
}

func (t *TagService) DeleteTag(ctx context.Context, userId string, tagId string) error {
	//validate that user have such tag
	if _, err := t.repo.GetByUserTagById(ctx, userId, tagId); err != nil {
		return err
	}

	return t.repo.DeleteTag(ctx, tagId)
}

func (t *TagService) UpdateTag(ctx context.Context, userId string, tagId string, input *model.TagInput) error {
	//validate that user have such tag
	if _, err := t.repo.GetByUserTagById(ctx, userId, tagId); err != nil {
		return err
	}

	// add id
	input.AddId(tagId)

	return t.repo.UpdateTag(ctx, input)
}

func (t *TagService) addUserIdToInput(tags []model.TagInput, userId string) {
	for _, it := range tags {
		it.AddUserId(userId)
	}
}
