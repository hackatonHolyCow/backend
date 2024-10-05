package items

import (
	"context"
	"hackathon/backend/entity"
	"hackathon/backend/internal/repository"
	"hackathon/backend/pkg/errors"
)

type ItemsService interface {
	List(ctx context.Context) ([]*entity.MenuItem, error)
}

type ItemsServiceImpl struct {
	repository *repository.Repository
}

func New(repo *repository.Repository) ItemsService {
	return &ItemsServiceImpl{
		repository: repo,
	}
}

func (i *ItemsServiceImpl) List(ctx context.Context) ([]*entity.MenuItem, error) {
	response, err := i.repository.Items.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "items: ItemsService.List repository.Items.List error")
	}

	return response, nil
}
