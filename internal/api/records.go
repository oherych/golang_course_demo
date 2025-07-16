package api

import (
	"github.com/labstack/echo/v4"
	"golang_course_demo/internal"
	"net/http"
	"strconv"
)

type RecordsController struct {
	recordsStorage recordsStorage
}

func (cc RecordsController) List(c echo.Context) error {
	items, err := cc.recordsStorage.All(c.Request().Context())
	if err != nil {
		return err
	}

	result := make([]map[string]any, len(items))
	for i, item := range items {
		result[i] = cc.toPresentation(item)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"items": result,
	})
}

func (cc RecordsController) ListByChannelID(c echo.Context) error {
	channelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": err.Error()})
	}

	items, err := cc.recordsStorage.ByChannelID(c.Request().Context(), channelID)
	if err != nil {
		return err
	}

	result := make([]map[string]any, len(items))
	for i, item := range items {
		result[i] = cc.toPresentation(item)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"items": result,
	})
}

func (cc RecordsController) toPresentation(item internal.Record) map[string]any {
	return map[string]any{
		"id":          item.ID,
		"title":       item.Title,
		"guid":        item.GUID,
		"description": item.Description,
		"link":        item.Link,
		"date":        item.PubDate,
	}
}
