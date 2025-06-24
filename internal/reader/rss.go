package reader

import (
	"context"
	"encoding/xml"
	"fmt"
	"golang_course_demo/internal"
	"net/http"
	"time"
)

type RSS struct {
	client http.Client
}

func New(client http.Client) RSS {
	return RSS{
		client: client,
	}
}

func (r RSS) Read(ctx context.Context, url string) (*internal.Channel, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	resp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d %s", resp.StatusCode, resp.Status)
	}

	d := xml.NewDecoder(resp.Body)

	var records Channel
	if err := d.Decode(&records); err != nil {
		return nil, fmt.Errorf("failed to decode RSS feed: %w", err)
	}

	return r.convert(records), nil
}

func (r RSS) convert(ch Channel) *internal.Channel {
	items := make([]internal.Record, len(ch.Items))
	for i, item := range ch.Items {
		items[i] = internal.Record{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			GUID:        item.GUID,
			PubDate:     time.Time(item.PubDate),
		}
	}

	return &internal.Channel{
		Title:       ch.Title,
		Description: ch.Description,
		Link:        ch.Link,
		Items:       items,
	}
}
