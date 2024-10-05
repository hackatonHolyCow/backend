package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateOrderItems, downCreateOrderItems)
}

func upCreateOrderItems(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS order_items (
			id text primary key not null,
			order_id uuid not null,
			item_id uuid not null,
			quantity int not null,
			price int not null,
			comments text
		);
	`)

	return err
}

func downCreateOrderItems(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `
		DROP TABLE order_items;
	`)

	return err
}
