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

func (t *TagService) GetAllTags(ctx context.Context) ([]*model.Tag, error) {
	return t.repo.GetAllTags(ctx)
}

func (t *TagService) CreateTags(ctx context.Context, tags *model.TagInput) ([]*model.Tag, error) {
	return t.repo.CreateTags(ctx, tags)
}

func (t *TagService) GetUserTags(ctx context.Context, userId string) ([]*model.Tag, error) {
	return t.repo.GetUserTags(ctx, userId)
}

func (t *TagService) DeleteTag(ctx context.Context, tagId string) error {
	return t.repo.DeleteTag(ctx, tagId)
}

func (t *TagService) UpdateTag(ctx context.Context, tagId string, input *model.TagInput) error {
	return t.repo.UpdateTag(ctx, tagId, input)
}
