package cronService

import (
	"fmt"
	"omt-project/services/timerService"
	"omt-project/utils"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func RunAllCronJobs(bot *linebot.Client) {
	s := New()
	s.Add(
		checkTimerAndPushQuizUrl(bot),
		1*time.Minute,
	)
	s.Run()
}

func checkTimerAndPushQuizUrl(bot *linebot.Client) func() {
	return func() {
		t := utils.JapanTimeNow()
		fmt.Println("cron doing good")
		currentTimerId := utils.GetTimerId(t)

		exist, err := timerService.ExistQuizTimer(currentTimerId)
		if err != nil {
			panic(err)
		}
		if !exist {
			return
		}

		timerRegisteredIds, err := timerService.ReadAllIdsByTimerId(currentTimerId)
		if err != nil {
			panic(err)
		}

		for _, userId := range timerRegisteredIds {
			message := linebot.NewTextMessage(utils.UserQuizUrl(userId))
			bot.PushMessage(userId, message).Do()
		}
	}
}
