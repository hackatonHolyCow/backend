package repository

import (
	"hackathon/backend/internal/repository/orders"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Orders orders.OrdersRepository
}

func New(psql *sqlx.DB) *Repository {
	return &Repository{
		Orders: orders.New(psql),
	}
}
