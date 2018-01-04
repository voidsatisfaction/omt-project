package quizService

import (
	"errors"
	"omt-project/services/wordService"
	"sort"
)

type QuizController struct {
	Words wordService.Words
}

func NewQuizController(userId string) (*QuizController, error) {
	wi, err := wordService.ReadWords(userId)
	if err != nil {
		return nil, err
	}

	qc := &QuizController{
		Words: wi.Words,
	}

	return qc, err
}

func (qc *QuizController) CreateMeaningWordQuiz(n int) (wordService.Words, error) {
	wordsSortedByPriority := qc.Words
	sort.Sort(byPriority(wordsSortedByPriority))
	if n == 0 {
		return nil, errors.New("CreateMeaningWordQuiz: n should be more than 0")
	}
	if len(wordsSortedByPriority) < n {
		return wordsSortedByPriority, nil
	}
	return wordsSortedByPriority[:n], nil
}

type byPriority wordService.Words

func (bp byPriority) Less(i, j int) bool {
	return bp[i].Priority > bp[j].Priority
}

func (bp byPriority) Len() int {
	return len(bp)
}

func (bp byPriority) Swap(i, j int) {
	bp[i], bp[j] = bp[j], bp[i]
}
