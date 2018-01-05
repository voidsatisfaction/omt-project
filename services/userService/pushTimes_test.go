package userService

import "testing"

func TestAddPushTimes(t *testing.T) {
	uid := "abc"
	timerId := "12:34"
	err := AddPushTimes(uid, timerId)
	if err != nil {
		t.Errorf("Add Push Times error! %+v", err)
	}
}
