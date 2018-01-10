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

func TestExistQuizTimer(t *testing.T) {
	tests := []struct {
		expect  bool
		timerId string
	}{
		{true, "2:34"},
		{false, "33:33"},
	}

	for _, test := range tests {
		expect, timerId := test.expect, test.timerId
		got, err := ExistQuizTimer(timerId)
		if err != nil {
			t.Errorf("ExistQuizTimer failed")
			t.Errorf("Error occured %+v", err)
		}

		if expect != got {
			t.Errorf("ExistQuizTimer failed")
			t.Errorf("Expect: %+v, got: %+v", expect, got)
		}
		fmt.Println(got)
	}
}
