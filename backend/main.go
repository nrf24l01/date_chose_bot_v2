package main

import (
	"dateChoice/handlers"
	"dateChoice/models"
	"dateChoice/routes"
	"dateChoice/schemas"

	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)
func main() {
	if os.Getenv("RUNTIME_PRODUCTION") != "true" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("failed to load .env: %v", err)
		}
	}

	db := models.RegisterPostgres()

	e := echo.New()


	if os.Getenv("RUNTIME_PRODUCTION") != "true" {
		e.Use(echoMw.Logger())
	}
    e.Use(echoMw.Recover())

	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{os.Getenv("ALLOWED_ORIGINS")},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	e.GET("/ping", func(c echo.Context) error {
	return c.JSON(200, schemas.Message{Status: "Date choice backend is ok"})
	})

	handler := &handlers.Handler{DB: db}
	routes.RegisterRoutes(e, handler)
	
	e.Logger.Fatal(e.Start(":1325"))
}