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

	group.GET("/choiced", h.GetUserDateChoiceHandler, middleware.TGMiddleware())
	group.GET("/all", h.GetAllUserChoiceHandler)
	group.GET("/interval", h.GetIntervalHandler)
}