package userService

import "testing"

func TestIsUserExist(t *testing.T) {
	userId := "abc"
	expect := true
	actual, err := IsUserExist(userId)
	if err != nil {
		t.Errorf("AWS S3 error, %+v", err)
	}

	if expect != actual {
		t.Errorf("User abc must be exist")
	}
}
