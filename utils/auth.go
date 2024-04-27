package utils

import (
	"os"

	"github.com/labstack/echo/v4"
)

func Authenticated(username, password string, c echo.Context) (bool, error) {
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if username == adminUsername && password == adminPassword {
		return true, nil
	}

	return false, nil
}
