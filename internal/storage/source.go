package storage

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang_course_demo/internal"
	"time"
)

var ErrNotFound = errors.New("not found")

type Source struct {
	db *pgxpool.Pool
}

func NewSource(db *pgxpool.Pool) Source {
	return Source{db: db}
}

func (r Source) Create(ctx context.Context, in internal.Source) (int, error) {
	row := r.db.QueryRow(ctx,
		"INSERT INTO sources (url, kind, created_at, updated_at) VALUES ($1, $2, $3, $4)  RETURNING id;",
		in.URL, in.Kind, time.Now(), time.Now(),
	)

	var id int
	return id, row.Scan(&id)
}

func (r Source) All(ctx context.Context) ([]internal.Source, error) {
	rows, err := r.db.Query(ctx, "SELECT id, url, kind, last_update_at, created_at, updated_at FROM sources")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sources []internal.Source
	for rows.Next() {
		var source internal.Source

		err = rows.Scan(&source.ID, &source.URL, &source.Kind, &source.LastUpdateAt, &source.CreatedAt, &source.UpdatedAt)
		if err != nil {
			return nil, err
		}

		sources = append(sources, source)
	}

	return sources, nil
}

func (r Source) Delete(ctx context.Context, id int) error {
	t, err := r.db.Exec(ctx, "DELETE FROM sources WHERE id = $1;", id)
	if err != nil {
		return err
	}

	if t.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
