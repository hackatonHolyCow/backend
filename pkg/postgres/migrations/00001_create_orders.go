package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateOrders, downCreateOrders)
}

func upCreateOrders(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS orders (
			id text primary key not null,
			total_amount bigint not null,
			board text not null,
			payment_id text not null
		);
	`)

	return err
}

func downCreateOrders(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `
		DROP TABLE orders;
	`)

	return err
}
