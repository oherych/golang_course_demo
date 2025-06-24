package worker

import (
	"context"
	"golang_course_demo/internal"
)

type Worker struct {
	reader readerInterface
	repo   repositoryInterface
}

type readerInterface interface {
	Read(ctx context.Context, url string) (*internal.Channel, error)
}

type repositoryInterface interface {
	CreateRecord(ctx context.Context, in internal.Record) error
}

func New(reader readerInterface, repo repositoryInterface) Worker {
	return Worker{
		reader: reader,
		repo:   repo,
	}
}

func (w Worker) Scan(ctx context.Context, urls []string) error {
	for _, url := range urls {
		channels, err := w.reader.Read(ctx, url)
		if err != nil {
			return err
		}

		for _, ch := range channels.Items {
			if err := w.repo.CreateRecord(ctx, ch); err != nil {
				return err
			}
		}

	}

	return nil
}
