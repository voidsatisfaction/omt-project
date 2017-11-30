package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// Setup
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	env := os.Getenv("APP_ENV")
	var port string
	switch env {
	case "PROD":
		port = os.Getenv("PORT")
	case "DEV":
		port = "9001"
	default:
		e.Logger.Fatal(`
      Please set up APP_ENV environment variable
      MODE=PROD or MODE=DEV
    `)
		return
	}

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		e.Logger.Fatal("There is a problem with making bot")
	}
	fmt.Printf("bot: %+v", bot)

	// Endpoints
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello medow")
	})

	e.POST("/bot/callback", func(c echo.Context) error {
		events, err := bot.ParseRequest(c.Request())
		fmt.Printf("events: %+v", events)
		if err != nil {
			e.Logger.Fatal(err)
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					e.Logger.Print(message)
					resMessage := linebot.NewTextMessage("こんにちは！")
					if _, err = bot.ReplyMessage(event.ReplyToken, resMessage).Do(); err != nil {
						e.Logger.Error(err)
					}
				}
			}
		}
		return c.JSON(http.StatusOK, "success")
	})

	// Listen
	e.Logger.Info(e.Start(fmt.Sprintf(":%s", port)))
}
