package orders

import (
	"context"
	"hackathon/backend/entity"
	"hackathon/backend/pkg/errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const ordersSchema = "orders"

type OrdersRepository interface {
	Create(ctx context.Context, order *entity.Order) (*entity.Order, error)
	Get(ctx context.Context, id string) (*entity.Order, error)
}

type OrdersRepositoryIplm struct {
	postgres *sqlx.DB
}

func New(psql *sqlx.DB) OrdersRepository {
	return &OrdersRepositoryIplm{
		postgres: psql,
	}
}

func (o *OrdersRepositoryIplm) Create(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	query, args, err := sq.
		Insert(ordersSchema).
		Columns("id", "total_amount", "board", "payment_id").
		Values(order.ID, order.TotalAmount, order.Table, order.PaymentID).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.WithHTTPCode(errors.Wrap(err, "orders: OrdersRepository.Create sq.ToSql error"), 500)
	}

	var response entity.Order
	if err := o.postgres.GetContext(ctx, &response, query, args...); err != nil {
		return nil, errors.WithHTTPCode(errors.Wrap(err, "orders: OrdersRepository.Create postgres.QueryRowxContext error"), 500)
	}

	return &response, nil
}

func (o *OrdersRepositoryIplm) Get(ctx context.Context, id string) (*entity.Order, error) {
	query, args, err := sq.
		Select("*").
		From(ordersSchema).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.WithHTTPCode(errors.Wrap(err, "orders: OrdersRepository.Get sq.ToSql error"), 500)
	}

	var response entity.Order
	if err := o.postgres.GetContext(ctx, &response, query, args...); err != nil {
		return nil, errors.WithHTTPCode(errors.Wrap(err, "orders: OrdersRepository.Get postgres.GetContext error"), 500)
	}

	return &response, nil
}
