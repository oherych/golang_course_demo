package worker

import (
	"context"
	"golang_course_demo/internal"
	"log"
	"time"
)

type Worker struct {
	reader   readerInterface
	channels channelsStorage
	repo     repositoryInterface
	source   sourceStorage
	ctx      context.Context
}

type readerInterface interface {
	Read(ctx context.Context, url string) (*internal.Channel, error)
}

type channelsStorage interface {
	Upset(ctx context.Context, in internal.Channel) (int, error)
}

type repositoryInterface interface {
	CreateRecord(ctx context.Context, channelID int, in internal.Record) error
}

type sourceStorage interface {
	All(ctx context.Context) ([]internal.Source, error)
}

func New(ctx context.Context, reader readerInterface, channels channelsStorage, repo repositoryInterface, source sourceStorage) Worker {
	return Worker{
		reader:   reader,
		channels: channels,
		repo:     repo,
		source:   source,
		ctx:      ctx,
	}
}

func (w Worker) Scan() error {
	for {
		select {
		case <-w.ctx.Done():
			return w.ctx.Err()
		case <-time.After(10 * time.Second):
			return w.scan()
		}
	}
}

func (w Worker) scan() error {
	log.Println("Scan started")

	sources, err := w.source.All(w.ctx)
	if err != nil {
		return err
	}

	for _, source := range sources {
		// TODO: kind

		channelData, err := w.reader.Read(w.ctx, source.URL)
		if err != nil {
			log.Println(err)
			continue
		}

		channel := internal.Channel{
			ID:          source.ID,
			Title:       channelData.Title,
			Description: channelData.Description,
			Link:        channelData.Link,
			Items:       channelData.Items,
		}

		channelID, err := w.channels.Upset(w.ctx, channel)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, ch := range channelData.Items {
			if err := w.repo.CreateRecord(w.ctx, channelID, ch); err != nil {
				log.Println(err)
				continue
			}
		}

	}

	return nil
}
