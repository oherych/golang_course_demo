package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang_course_demo/internal"
)

type Router struct {
	address string
	e       *echo.Echo
}

type recordsStorage interface {
	All(ctx context.Context) ([]internal.Record, error)
	ByChannelID(ctx context.Context, channelID int) ([]internal.Record, error)
}

type channelsStorage interface {
	All(ctx context.Context) ([]internal.Channel, error)
}

type sourceStorage interface {
	Create(ctx context.Context, in internal.Source) (int, error)
	All(ctx context.Context) ([]internal.Source, error)
	Delete(ctx context.Context, id int) error
}

func New(address string, source sourceStorage, channels channelsStorage, records recordsStorage) Router {
	sourceController := SourceController{sourceStorage: source}
	channelsController := ChannelsController{channelsStorage: channels}
	recordsController := RecordsController{recordsStorage: records}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())

	e.GET("/source/", sourceController.List)
	e.POST("/source/", sourceController.Create)
	e.DELETE("/source/:id/", sourceController.Delete)

	e.GET("/channels/", channelsController.List)
	e.GET("/channels/:id/records/", recordsController.ListByChannelID)

	e.GET("/records/", recordsController.List)

	return Router{
		address: address,
		e:       e,
	}
}

func (r *Router) Start() error {
	return r.e.Start(r.address)
}
