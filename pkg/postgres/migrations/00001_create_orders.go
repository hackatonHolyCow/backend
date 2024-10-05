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
			id uuid primary key default gen_random_uuid(),
			total_amount double_precission not null,
			board text not null
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
