package orders

import (
	"context"
	"hackathon/backend/entity"
	"hackathon/backend/internal/repository"
	"hackathon/backend/pkg/errors"
)

type OrdersService interface {
	Create(ctx context.Context, order *entity.Order) (*entity.Order, error)
}

type OrdersServiceImpl struct {
	repository *repository.Repository
}

func New(repo *repository.Repository) OrdersService {
	return &OrdersServiceImpl{
		repository: repo,
	}
}

func (o *OrdersServiceImpl) Create(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	response, err := o.repository.Orders.Create(ctx, order)
	if err != nil {
		return nil, errors.Wrap(err, "orders: OrdersService.Create repository.Orders.Create error")
	}

	return response, nil
}
