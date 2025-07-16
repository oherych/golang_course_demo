package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang_course_demo/internal"
)

type Channel struct {
	db *pgxpool.Pool
}

func NewChannel(db *pgxpool.Pool) Channel {
	return Channel{db: db}
}

func (r Channel) Upset(ctx context.Context, in internal.Channel) (int, error) {
	row := r.db.QueryRow(ctx,
		"INSERT INTO channel (id, title, description, link) VALUES ($1, $2, $3, $4) ON CONFLICT (id) DO UPDATE SET title = EXCLUDED.title, description = EXCLUDED.description, link = EXCLUDED.link RETURNING id;",
		in.ID, in.Title, in.Description, in.Link,
	)

	var id int
	return id, row.Scan(&id)
}

func (r Channel) All(ctx context.Context) ([]internal.Channel, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, description, link FROM channel")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []internal.Channel
	for rows.Next() {
		var item internal.Channel

		err = rows.Scan(&item.ID, &item.Title, &item.Description, &item.Link)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
