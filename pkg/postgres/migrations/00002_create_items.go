package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateItems, downCreateItems)
}

func upCreateItems(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS items (
			id uuid primary key default gen_random_uuid(),
			name text not null,
			description text not null,
			price bigint not null,
			tags text[] not null,
			picture text not null
		);
	`)

	return err
}

func downCreateItems(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `
		DROP TABLE items
	`)

	return err
}
