package middleware

import (
	"dateChoice/schemas"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func TGMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, schemas.Error{Error: "missing token"})
			}

			// Убираем "Bearer " из заголовка
			tokenString := ""
			if len(authHeader) > 4 && authHeader[:4] == "tma " {
				tokenString = authHeader[4:]
			} else {
				return c.JSON(http.StatusUnauthorized, schemas.Error{Error: "invalid token format"})
			}

			expInStr := os.Getenv("INIT_DATA_EXPIRE_HOURS")
			expIn, err := strconv.Atoi(expInStr)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, schemas.Error{Error: "Server error"})
			}
			expInHour := time.Duration(expIn) * time.Hour
			
			verifyErr := initdata.Validate(tokenString, os.Getenv("BOT_TOKEN"), expInHour)
			if verifyErr != nil {
				return c.JSON(http.StatusUnauthorized, schemas.Error{Error: "invalid token"})
			}
			tokenData, err := initdata.Parse(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, schemas.Error{Error: "invalid token"})
			}
			if tokenData.User.ID != 0 {
				c.Set("userID", tokenData.User.ID)
				c.Set("userName", tokenData.User.Username)
				log.Printf("User ID: %d, Username: %s", tokenData.User.ID, tokenData.User.Username)
			} else {
				return c.JSON(http.StatusUnauthorized, schemas.Error{Error: "token does not contain user data"})
			}
			return next(c)
		}
	}
}