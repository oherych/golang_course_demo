package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang_course_demo/internal"
)

type Feeds struct {
	db *pgxpool.Pool
}

func NewFeeds(db *pgxpool.Pool) Feeds {
	return Feeds{db: db}
}

func (r Feeds) CreateRecord(ctx context.Context, in internal.Record) error {
	_, err := r.db.Exec(ctx,
		"INSERT INTO records (guid, title, link, description, pub_date) VALUES ($1, $2, $3, $4, $5)",
		in.GUID, in.Title, in.Link, in.Description, in.PubDate,
	)

	return err
}
