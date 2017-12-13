package botHandler

import (
	"fmt"
	"net/http"
	"omt-project/services/botService"
	"omt-project/services/userService"
	"omt-project/services/wordService"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

// TODO: refer https://github.com/line/line-bot-sdk-go/blob/master/examples/kitchensink/server.go
func CallbackHandlerGenerator(e *echo.Echo, bot *linebot.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			e.Logger.Fatal(err)
		}

		for _, event := range events {
			fmt.Printf("event: %+v\n", event)
			fmt.Printf("userId: %+v\n", event.Source.UserID)
			switch event.Type {
			case linebot.EventTypeFollow:
				uid := event.Source.UserID
				go userService.CreateNewUser(uid)
				go wordService.CreateNewWord(uid)
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					uid := event.Source.UserID
					action := botService.CreateAction(uid, message, event.ReplyToken, event.Source)
					actionResult := botService.TreatAction(action)
					e.Logger.Printf("message: %+v\n", message)
					e.Logger.Printf("action: %+v\n", action)
					e.Logger.Printf("actionResult: %+v\n", actionResult)
					resMessage := linebot.NewTextMessage(actionResult.Text)
					if _, err = bot.ReplyMessage(event.ReplyToken, resMessage).Do(); err != nil {
						e.Logger.Error(err)
					}
				}
			default:

			}
		}
		return c.JSON(http.StatusOK, "success")
	}
}
