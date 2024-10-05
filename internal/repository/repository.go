package repository

import (
	"hackathon/backend/internal/repository/items"
	"hackathon/backend/internal/repository/orders"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Orders orders.OrdersRepository
	Items  items.ItemsRepository
}

func New(psql *sqlx.DB) *Repository {
	return &Repository{
		Orders: orders.New(psql),
		Items:  items.New(psql),
	}
}
