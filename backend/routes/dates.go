package routes

import (
	"dateChoice/handlers"
	"dateChoice/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterDateRoutes(e *echo.Echo, h* handlers.Handler) {
	group := e.Group("/date")

	group.POST("/choice", h.DateChoiceHandler, middleware.TGMiddleware())
}