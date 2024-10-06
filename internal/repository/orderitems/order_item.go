package orderitems

import (
	"context"
	"hackathon/backend/entity"
	"hackathon/backend/pkg/errors"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
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
		Columns("id", "order_id", "item_id", "quantity", "comments", "price").
		Values(uuid.NewString(), order.OrderID, order.ItemID, order.Quantity, order.Comments, order.Price).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.WithHTTPCode(errors.Wrap(err, "orderitems: OrderItemsRepository.Create sq.ToSql error"), http.StatusInternalServerError)
	}

	var response entity.OrderItem
	if err := o.postgres.GetContext(ctx, &response, query, args...); err != nil {
		return nil, errors.WithHTTPCode(errors.Wrap(err, "orderitems: OrderItemsRepository.Create postgres.GetContext error"), http.StatusInternalServerError)
	}

	return &response, nil
}
