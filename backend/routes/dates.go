package routes

import (
	"dateChoice/handlers"
	"dateChoice/middleware"
	"dateChoice/schemas"

	"github.com/labstack/echo/v4"
)

func RegisterDateRoutes(e *echo.Echo, h* handlers.Handler) {
	group := e.Group("/date")

	group.POST("/choice", h.DateChoiceHandler, middleware.TGMiddleware(), middleware.ValidationMiddleware(func() interface{} {
		return &schemas.DateChoiceRequest{}
	}))
}