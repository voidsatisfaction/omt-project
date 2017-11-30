package botHandler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

func CallbackHandlerGenerator(e *echo.Echo, bot *linebot.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
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
	}
}
