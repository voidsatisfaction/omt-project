package main

import (
	"fmt"
	"net/http"

	"omt-project/config"
	"omt-project/handlers/botHandler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// Setup
	e := echo.New()
	cfg := config.Setting()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	bot, err := linebot.New(
		cfg.ChannelSecret,
		cfg.ChannelToken,
	)
	if err != nil {
		e.Logger.Fatal("There is a problem with making bot")
	}

	// Endpoints
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Worked")
	})

	e.POST("/bot/callback", botHandler.CallbackHandlerGenerator(e, bot))

	// Listen
	e.Logger.Info(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
