package items

import (
	"context"
	"fmt"
	"hackathon/backend/entity"
	"hackathon/backend/pkg/errors"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const itemsSchema = "items"

type ItemsRepository interface {
	List(ctx context.Context) ([]*entity.MenuItem, error)
	ListByOrderID(ctx context.Context, orderID string) ([]*entity.MenuItem, error)
}

type ItemsRepositoryImpl struct {
	postgres *sqlx.DB
}

func New(postgres *sqlx.DB) ItemsRepository {
	return &ItemsRepositoryImpl{
		postgres: postgres,
	}
}

func (i *ItemsRepositoryImpl) List(ctx context.Context) ([]*entity.MenuItem, error) {
	query, args, err := sq.
		Select("*").
		From(itemsSchema).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.WithHTTPCode(errors.Wrap(err, "items: ItemsRepository.List sq.ToSql error"), http.StatusInternalServerError)
	}

	response := make([]*entity.MenuItem, 0)
	if err := i.postgres.Select(&response, query, args...); err != nil {
		return nil, errors.WithHTTPCode(errors.Wrap(err, "items: ItemsRepository.List postgres.Select error"), http.StatusInternalServerError)
	}

	return response, nil
}

func (i *ItemsRepositoryImpl) ListByOrderID(ctx context.Context, orderID string) ([]*entity.MenuItem, error) {
	query, args, err := sq.Select("i.*", "oi.comments", "oi.quantity").
		From(fmt.Sprintf("%s i", itemsSchema)).
		Join("order_items oi ON i.id = oi.item_id").
		Where(sq.Eq{"oi.order_id": orderID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.WithHTTPCode(errors.Wrap(err, "items: ItemsRepository.ListByOrderID sq.ToSql error"), http.StatusInternalServerError)
	}

	response := make([]*entity.MenuItem, 0)
	if err := i.postgres.Select(&response, query, args...); err != nil {
		return nil, errors.WithHTTPCode(errors.Wrap(err, "items: ItemsRepository.ListByOrderID postgres.Select error"), http.StatusInternalServerError)
	}

	return response, nil
}
