package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	env := os.Getenv("APP_ENV")
	var port string
	switch env {
	case "PROD":
		port = "9000"
	case "DEV":
		port = "9001"
	default:
		e.Logger.Fatal(`
      Please set up APP_ENV environment variable
      MODE=PROD or MODE=DEV
    `)
		return
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello medow")
	})
	e.Logger.Info(e.Start(fmt.Sprintf(":%s", port)))
}
