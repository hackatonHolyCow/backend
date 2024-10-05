package orderitems

import (
	"context"
	"hackathon/backend/entity"
	"hackathon/backend/pkg/errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const orderItemsSchema = "order_items"

type OrderItemsRepository interface {
	Create(ctx context.Context, order *entity.OrderItem) (*entity.OrderItem, error)
}

type OrderItemsRepositoryImpl struct {
	postgres *sqlx.DB
}

func New(postgres *sqlx.DB) OrderItemsRepository {
	return &OrderItemsRepositoryImpl{
		postgres: postgres,
	}
}

func (o *OrderItemsRepositoryImpl) Create(ctx context.Context, order *entity.OrderItem) (*entity.OrderItem, error) {
	query, args, err := sq.
		Insert(orderItemsSchema).
		Columns("order_id", "item_id", "quantity", "comments").
		Values(order.OrderID, order.ItemID, order.Quantity, order.Comments).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "orderitems: OrderItemsRepository.Create error")
	}

	var response entity.OrderItem
	if err := o.postgres.GetContext(ctx, &response, query, args...); err != nil {
		return nil, errors.Wrap(err, "orderitems: OrderItemsRepository.Create error")
	}

	return &response, nil
}
