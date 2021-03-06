package main

import (
	"fmt"
	"net/http"

	"omt-project/config"
	"omt-project/src/handlers/botHandler"
	"omt-project/src/handlers/webHandler"
	"omt-project/src/services/cronService"
	"omt-project/src/templateEngine"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// Setup
	e := echo.New()
	e.Renderer = templateEngine.NewHtmlTemplateEngine()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	cfg := config.Setting()
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

	e.POST("/quiz", webHandler.PostQuizHandlerGenerator(e))
	e.GET("/quiz/new/:userId", webHandler.GetQuizHandlerGenerator(e))
	e.POST("/quiz/result", webHandler.GetQuizResultHandlerGenerator(e))

	e.POST("/bot/callback", botHandler.CallbackHandlerGenerator(e, bot))

	// Cron jobs
	cronService.RunAllCronJobs(bot)

	// Listen
	e.Logger.Info(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
