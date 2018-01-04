package timerService

import (
	"fmt"
	"testing"
)

func TestReadQuizTimer(t *testing.T) {
	timerId := "2:34"
	qtm, _ := ReadQuizTimer(timerId)
	fmt.Printf("%+v\n", qtm)
}
