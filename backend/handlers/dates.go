package handlers

import (
	"dateChoice/models"
	"dateChoice/schemas"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *Handler) DateChoiceHandler(c echo.Context) error {
	dates := c.Get("validatedBody").(*schemas.DateChoiceRequest)

	var userChoice models.UserChoice
	if err := h.DB.First(&userChoice, "user_id = ?", c.Get("userID")).Error; err != nil {
		if err.Error() == "record not found" {
			userChoice.UserID = c.Get("userID").(int64)
			userChoice.UserName = c.Get("userName").(string)
		} else {
			return c.JSON(http.StatusInternalServerError, schemas.Error{Error: "Server error"})
		}
	}
	userChoice.SelectedDates = pq.StringArray(dates.Dates)
	if err := h.DB.Save(&userChoice).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, schemas.Error{Error: "Server error"})
	}
	return c.JSON(http.StatusOK, schemas.Message{Status: "Dates saved successfully"})
}