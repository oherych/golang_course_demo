package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang_course_demo/internal"
)

//

type Records struct {
	db *pgxpool.Pool
}

func NewRecords(db *pgxpool.Pool) Records {
	return Records{db: db}
}

func (r Records) CreateRecord(ctx context.Context, channelID int, in internal.Record) error {
	_, err := r.db.Exec(ctx,
		"INSERT INTO records (channel_id, guid, title, link, description, pub_date) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (guid) DO NOTHING",
		channelID, in.GUID, in.Title, in.Link, in.Description, in.PubDate,
	)

	return err
}

func (r Records) All(ctx context.Context) ([]internal.Record, error) {
	rows, err := r.db.Query(ctx, "SELECT id, guid, title, link, description, pub_date FROM records")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return r.scanList(rows)
}

func (r Records) ByChannelID(ctx context.Context, channelID int) ([]internal.Record, error) {
	rows, err := r.db.Query(ctx, "SELECT id, guid, title, link, description, pub_date FROM records WHERE channel_id = $1", channelID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return r.scanList(rows)
}

func (r Records) scanList(rows pgx.Rows) ([]internal.Record, error) {
	var sources []internal.Record
	for rows.Next() {
		var source internal.Record

		err := rows.Scan(&source.ID, &source.GUID, &source.Title, &source.Link, &source.Description, &source.PubDate)
		if err != nil {
			return nil, err
		}

		sources = append(sources, source)
	}

	return sources, nil
}
