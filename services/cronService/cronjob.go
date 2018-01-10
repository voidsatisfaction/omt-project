package cronService

import (
	"fmt"
	"omt-project/services/timerService"
	"omt-project/utils"
	"time"
)

func RunAllCronJobs() {
	s := New()
	s.Add(checkTimerAndPushQuizUrl, 1*time.Minute)
	s.Run()
}

func checkTimerAndPushQuizUrl() {
	currentTimerId := utils.GetTimerId(time.Now())

	exist, err := timerService.ExistQuizTimer(currentTimerId)
	if err != nil {
		panic(err)
	}
	if !exist {
		return
	}

	quizTimerMap, err := timerService.ReadQuizTimer(currentTimerId)
	if err != nil {
		panic(err)
	}
	fmt.Println(quizTimerMap)
}
