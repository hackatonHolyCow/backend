package service

import (
	"hackathon/backend/internal/repository"
	"hackathon/backend/internal/service/items"
	"hackathon/backend/internal/service/orders"
)

type Service struct {
	Orders orders.OrdersService
	Items  items.ItemsService
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Orders: orders.New(repo),
		Items:  items.New(repo),
	}
}
