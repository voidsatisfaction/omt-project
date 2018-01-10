package utils

import (
	"fmt"
	"time"
)

func GetTimerId(current time.Time) string {
	currentHour, currentMinute := current.Hour(), current.Minute()
	if currentMinute < 10 {
		return fmt.Sprintf("%d:0%d", currentHour, currentMinute)
	}
	return fmt.Sprintf("%d:%d", currentHour, currentMinute)
}
