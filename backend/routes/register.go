package routes

import (
	"dateChoice/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, h *handlers.Handler) {
	RegisterDateRoutes(e, h)
}
