package wordService

import "testing"

func TestCreateNewWord(t *testing.T) {
	userId := "abc"
	CreateNewWord(userId)
}

func TestReadWord(t *testing.T) {
	userId := "abc"
	ReadWord(userId)
}

func TestAddWord(t *testing.T) {
	userId := "abc"
	wordName := "water"
	meaning := []string{"水", "みず"}
	err := Addword(userId, wordName, meaning)
	if err != nil {
		t.Errorf("There was error: %+v", err)
	}
}
