package service

import (
	"hackathon/backend/internal/repository"
	"hackathon/backend/internal/service/items"
	"hackathon/backend/internal/service/orders"

	"github.com/mercadopago/sdk-go/pkg/config"
)

type Service struct {
	Orders orders.OrdersService
	Items  items.ItemsService
}

func New(repo *repository.Repository, mpConfig config.Config) *Service {
	return &Service{
		Orders: orders.New(repo, mpConfig),
		Items:  items.New(repo),
	}
}
