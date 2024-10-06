package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddOrderStatus, downAddOrderStatus)
}

func upAddOrderStatus(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, `
		ALTER TABLE orders ADD COLUMN status text NOT NULL DEFAULT 'pending';
	`)

	return err
}

func downAddOrderStatus(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `
		ALTER TABLE orders DROP COLUMN status;
	`)

	return err
}
