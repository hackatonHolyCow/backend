package service

import (
	"hackathon/backend/internal/repository"
	"hackathon/backend/internal/service/orders"
)

type Service struct {
	Orders orders.OrdersService
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Orders: orders.New(repo),
	}
}
