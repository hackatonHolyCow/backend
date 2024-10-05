package items

import (
	"context"
	"hackathon/backend/entity"
	"hackathon/backend/pkg/errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const itemsSchema = "items"

type ItemsRepository interface {
	List(ctx context.Context) ([]*entity.MenuItem, error)
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
		return nil, errors.Wrap(err, "items: ItemsRepository.List sq.ToSql errro")
	}

	response := make([]*entity.MenuItem, 0)
	if err := i.postgres.Select(&response, query, args...); err != nil {
		return nil, errors.Wrap(err, "items: ItemsRepository.List postgres.Select errro")
	}

	return response, nil
}
