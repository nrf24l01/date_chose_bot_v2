package handlers

import (
	"dateChoice/models"
	"dateChoice/schemas"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *Handler) DateChoiceHandler(c echo.Context) error {
	dates := c.Get("validatedBody").(*schemas.DateChoiceRequest)

	allowFrom := os.Getenv("ALLOW_FROM")
	allowTo := os.Getenv("ALLOW_TO")

	layout := "2006-01-02"
	allowFromTime, err := time.Parse(layout, allowFrom)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schemas.Error{Error: "Invalid ALLOW_FROM format"})
	}
	allowToTime, err := time.Parse(layout, allowTo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schemas.Error{Error: "Invalid ALLOW_TO format"})
	}

	if len(dates.Dates) != 1 || dates.Dates[0] != "0001-01-01" {
		for _, dateStr := range dates.Dates {
			dateTime, err := time.Parse(layout, dateStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, schemas.Error{Error: "Invalid date format (YYYY-MM-DD)"})
			}
			if dateTime.Before(allowFromTime) || dateTime.After(allowToTime) {
				return c.JSON(http.StatusBadRequest, schemas.Error{Error: "Date out of allowed range (YYYY-MM-DD)"})
			}
		}
	}

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

func (h *Handler) GetUserDateChoiceHandler(c echo.Context) error {
	var userChoice models.UserChoice
	if err := h.DB.First(&userChoice, "user_id = ?", c.Get("userID")).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusOK, schemas.UserDateChoiceResponse{Dates: []string{}})
		}
		return c.JSON(http.StatusInternalServerError, schemas.Error{Error: "Server error"})
	}

	return c.JSON(http.StatusOK, schemas.UserDateChoiceResponse{Dates: userChoice.SelectedDates})
}