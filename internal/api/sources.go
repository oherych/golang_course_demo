package api

import (
	"github.com/labstack/echo/v4"
	"golang_course_demo/internal"
	"golang_course_demo/internal/storage"
	"net/http"
	"strconv"
)

type SourceController struct {
	sourceStorage sourceStorage
}

func (cc SourceController) Create(c echo.Context) error {
	var target struct {
		URL  string `json:"url"`
		Kind string `json:"kind"`
	}
	if err := c.Bind(&target); err != nil {
		return err
	}

	id, err := cc.sourceStorage.Create(c.Request().Context(), internal.Source{
		URL:  target.URL,
		Kind: target.Kind,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]int{"id": id})
}

func (cc SourceController) List(c echo.Context) error {
	items, err := cc.sourceStorage.All(c.Request().Context())
	if err != nil {
		return err
	}

	result := make([]map[string]any, len(items))
	for i, item := range items {
		result[i] = map[string]any{
			"id":             item.ID,
			"kind":           item.Kind,
			"url":            item.URL,
			"last_update_at": item.LastUpdateAt,
			"created_at":     item.CreatedAt,
			"updated_at":     item.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, map[string]any{
		"items": result,
	})
}

func (cc SourceController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": err.Error()})
	}

	err = cc.sourceStorage.Delete(c.Request().Context(), id)
	if err == storage.ErrNotFound {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
