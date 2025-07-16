package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ChannelsController struct {
	channelsStorage channelsStorage
}

func (cc ChannelsController) List(c echo.Context) error {
	items, err := cc.channelsStorage.All(c.Request().Context())
	if err != nil {
		return err
	}

	result := make([]map[string]any, len(items))
	for i, item := range items {
		result[i] = map[string]any{
			"id":          item.ID,
			"title":       item.Title,
			"description": item.Description,
			"link":        item.Link,
		}
	}

	return c.JSON(http.StatusOK, map[string]any{
		"items": result,
	})
}
