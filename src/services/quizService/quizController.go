package quizService

import (
	"errors"
	"omt-project/src/services/wordService"
	"omt-project/src/utils"
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
	// if quiz numbers are more than words
	// it returns all words as quiz
	if len(wordsSortedByPriority) <= n {
		return wordsSortedByPriority, nil
	}

	// if former 0 ~ k index are all same value(k > n),
	// computer randomize these
	lastSamePriorityIndex := 0
	for i, word := range wordsSortedByPriority {
		if i == len(wordsSortedByPriority)-1 {
			break
		}
		nextWord := wordsSortedByPriority[i+1]
		if word.Priority > nextWord.Priority {
			break
		}
		lastSamePriorityIndex++
	}
	if lastSamePriorityIndex > n {
		wordsUntilSamePriority := wordsSortedByPriority[:lastSamePriorityIndex+1]
		utils.Shuffle(byPriority(wordsUntilSamePriority))
		return wordsUntilSamePriority[:n], nil
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

func (qc *QuizController) ApplyMeaningWordQuizResult(srs ScoreResults) {
	qc.updatetWordPriority(srs)
}

func (qc *QuizController) updatetWordPriority(srs ScoreResults) {
	oldWords := qc.Words
	for _, sr := range srs {
		// TODO: It is O(n*m)
		newWordName := sr.Word
		correctNess := sr.Correctness
		for i, oldWord := range oldWords {
			oldWordName := oldWord.Name
			if oldWordName == newWordName {
				oldWordPriority := oldWord.Priority
				if correctNess {
					oldWords[i].Priority = calculateCorrectPriority(oldWordPriority)
				} else {
					oldWords[i].Priority = calculateNoCorrectPriority(oldWordPriority)
				}
				break
			}
		}
	}
	qc.Words = oldWords
}

func calculateCorrectPriority(priority int) int {
	newPriority := priority / 2
	return newPriority
}

func calculateNoCorrectPriority(priority int) int {
	// 0 <= priority <= 100
	newPriority := (priority + 100) / 2
	return newPriority
}

type ScoreResults []*ScoreResult

type ScoreResult struct {
	Word        string
	Correctness bool
}

func GetScoreResults(userAns []string, goodAns []string) ScoreResults {
	n := len(userAns)
	srs := make(ScoreResults, 0, n)
	for i := 0; i < n; i++ {
		if userAns[i] == goodAns[i] {
			sr := &ScoreResult{goodAns[i], true}
			srs = append(srs, sr)
		} else {
			sr := &ScoreResult{goodAns[i], false}
			srs = append(srs, sr)
		}
	}
	return srs
}

func (qc *QuizController) UpdateWordsInfo(userId string) error {
	words := qc.Words
	wi := wordService.NewWordsInfo()
	wi.AssignWords(words)

	err := wordService.UpdateNewWord(userId, wi)
	if err != nil {
		return err
	}
	return nil
}
