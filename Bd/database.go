package bd

import (
	"context"
	datework "micron/DateWork"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// connection to localhost
func CreateConnection(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, "postgres://anme:anme228@localhost:5432/telegram")
}

// Create task database
func CreateDatabase(ctx context.Context, pool *pgxpool.Pool) error {
	sqlQuery := `
		CREATE  TABLE IF NOT EXISTS timers(
			id SERIAL PRIMARY KEY,
			title VARCHAR(200),
			command VARCHAR(200),
			args TEXT[],
			time_implementation TIMESTAMP NOT NULL,
			repeating BOOLEAN NOT NULL DEFAULT FALSE,
			done BOOLEAN NOT NULL DEFAULT FALSE,
			time_created TIMESTAMP

		)
	`
	if _, err := pool.Exec(ctx, sqlQuery); err != nil {
		return err
	}
	return nil
}

// Write task to db
func WriteTimer(ctx context.Context, pool *pgxpool.Pool, d *datework.DateJson) error {
	sqlQuery := `
		INSERT INTO timers(title, command, args, time_implementation, repeating, time_created)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	if err := pool.QueryRow(ctx, sqlQuery, d.Description, d.Command, d.Args, d.Date, d.Repeat, time.Now()).Scan(&d.ID); err != nil {
		return err
	}
	return nil
}

// Set task update done
func UpdateTaskDone(d datework.DateJson, ctx context.Context, p *pgxpool.Pool) error {
	sqlQuery := `
	UPDATE timers SET done = true
	WHERE id = $1;
	`
	if _, err := p.Exec(ctx, sqlQuery, d.ID); err != nil {
		return err
	}
	return nil
}

// Update task in db
func UpdateTask(d datework.DateJson, ctx context.Context, p *pgxpool.Pool) error {
	sqlQuery := `
		UPDATE timers SET
			title = $1,
			command = $2,
			time_implementation = $3,
			repeating = $4,
			args = $5
		WHERE id = $6
	`
	if _, err := p.Exec(ctx, sqlQuery, d.Description, d.Command, d.Date, d.Repeat, d.Args, d.ID); err != nil {
		return err
	}
	return nil
}

func CheckDbNotDone(ctx context.Context, p *pgxpool.Pool) ([]datework.DateJson, error) {
	sqlQuery := `
		SELECT id, title, command, time_implementation, repeating, time_created, args, done
		FROM timers
		WHERE done = false
		ORDER BY id ASK; 
	`
	rows, err := p.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Dates := make([]datework.DateJson, 1)
	for rows.Next() {
		var temp datework.DateJson
		if err := rows.Scan(
			&temp.ID,
			&temp.Description,
			&temp.Command,
			&temp.Date,
			&temp.Repeat,
			&temp.TimeCreated,
			&temp.Args,
			&temp.Done,
		); err != nil {
			return nil, err
		}
		Dates = append(Dates, temp)
	}

	return Dates, nil
}

func CheckFullDb(ctx context.Context, p *pgxpool.Pool) ([]datework.DateJson, error) {
	sqlQuery := `
		SELECT id, title, command, time_implementation, repeating, time_created, args, done
		FROM timers
		ORDER BY id ASC;
	`
	rows, err := p.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Dates := make([]datework.DateJson, 1)
	for rows.Next() {
		var temp datework.DateJson
		if err := rows.Scan(
			&temp.ID,
			&temp.Description,
			&temp.Command,
			&temp.Date,
			&temp.Repeat,
			&temp.TimeCreated,
			&temp.Args,
			&temp.Done,
		); err != nil {
			return nil, err
		}
		Dates = append(Dates, temp)
	}

	return Dates, nil
}

func DeleteTaskById(ctx context.Context, p *pgxpool.Pool, id int) error {
	sqlQuery := `
	DELETE FROM timers
	WHERE id = $1
	`
	if _, err := p.Exec(ctx, sqlQuery, id); err != nil {
		return err
	}
	return nil
}
