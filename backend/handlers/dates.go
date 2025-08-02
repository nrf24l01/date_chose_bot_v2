package handlers

import (
	"dateChoice/schemas"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DateChoiceHandler(c echo.Context) error {
	return c.JSON(http.StatusCreated, schemas.Message{Status: "Date choice"})
}