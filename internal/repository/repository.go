package repository

import (
	"hackathon/backend/internal/repository/items"
	"hackathon/backend/internal/repository/orderitems"
	"hackathon/backend/internal/repository/orders"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Orders     orders.OrdersRepository
	Items      items.ItemsRepository
	OrderItems orderitems.OrderItemsRepository
}

func New(psql *sqlx.DB) *Repository {
	return &Repository{
		Orders:     orders.New(psql),
		Items:      items.New(psql),
		OrderItems: orderitems.New(psql),
	}
}
