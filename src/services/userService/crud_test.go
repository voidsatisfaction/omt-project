package userService

import (
	"testing"
)

func TestCreateNewUser(t *testing.T) {
	userId := "abc"
	var expect error = nil
	err := CreateNewUser(userId)
	if err != expect {
		t.Errorf("Expect: %+v, got: %+v", expect, err)
	}
}

func TestReadUserInfo(t *testing.T) {
	userId := "abc"
	var expect error = nil
	_, err := ReadUserInfo(userId)
	if err != expect {
		t.Errorf("Expect: %+v, got: %+v", expect, err)
	}
}
