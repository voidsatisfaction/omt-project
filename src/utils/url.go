package utils

import (
	"fmt"
	"omt-project/config"
)

func UserQuizUrl(userId string) string {
	c := config.Setting()
	return fmt.Sprintf("%s/quiz/new/%s", c.Host, userId)
}
