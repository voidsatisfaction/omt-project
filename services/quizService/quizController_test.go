package quizService

import (
	"testing"
)

func TestNewQuizController(t *testing.T) {
	userId := "abc"
	_, err := NewQuizController(userId)
	if err != nil {
		t.Errorf("NewQuizController error %+v\n", err)
	}
}

func TestCreateMeaningWordQuiz(t *testing.T) {
	userId := "abc"
	qc, _ := NewQuizController(userId)
	_, err := qc.CreateMeaningWordQuiz(0)
	if err == nil {
		t.Errorf("Word Quiz should be more than 0")
	}
	wq2, err := qc.CreateMeaningWordQuiz(10000)
	if err != nil {
		t.Errorf("Create Meaning Word Quiz error: %+v\n", err)
	}

	if len(wq2) != len(qc.Words) {
		t.Errorf("word quiz should be itself if number is too huge")
	}
}
