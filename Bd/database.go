package bd

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateDatabase(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
		CREATE  TABLE IF NOT EXISTS timers(
			id SERIAL PRIMARY KEY,
			title VARCHAR(200),
			command VARCHAR(200),
			time_implementation TIMESTAMP NOT NULL,
			repeating BOOLEAN NOT NULL,
			time_created TIMESTAMP

		)
	`
	if _, err := conn.Exec(ctx, sqlQuery); err != nil {
		return err
	}
	return nil
}

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:anme228@localhost:5432/postgres")
}
