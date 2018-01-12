package botService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"omt-project/config"
	"omt-project/src/api/glosbe"
	"omt-project/src/services/timerService"
	"omt-project/src/services/userService"
	"omt-project/src/services/wordService"
	"strings"
)

const (
	SuccessCode ActionStatusCode = "SUCCESS"
	FailCode    ActionStatusCode = "FAIL"
)

type ActionStatusCode string

type ActionResult struct {
	Status ActionStatusCode
	Text   string
}

func newActionResult() *ActionResult {
	return &ActionResult{
		Status: SuccessCode,
	}
}

func (ar *ActionResult) ServerError() *ActionResult {
	ar.Text = "先生、今、体の調子が悪いの"
	ar.Status = FailCode
	return ar
}

func (ar *ActionResult) PhraseNotFound() *ActionResult {
	bp := BotPhrase{}
	bp.Setting()
	ar.Text = string(bp[PhraseNotFound])
	ar.Status = FailCode
	return ar
}

// TreatAction is main treater of actions
func TreatAction(a *Action) *ActionResult {
	actionResult := newActionResult()
	switch a.ActionType {
	case Invalid:
		actionResult = TreatPredefinedAction(a)
	case InvalidCommand:
		actionResult = TreatPredefinedAction(a)
	case Search:
		// Call word search api
		actionResult = TreatSearchAction(a)
	case Add:
		// Add user own phrase
		actionResult = TreatAddAction(a)
	case All:
		actionResult = TreatAllAction(a)
	case Set:
		actionResult = TreatSetAction(a)
	case TimerAll:
		actionResult = TreatTimerAllAction(a)
	case Quiz:
		actionResult = TreatQuizAction(a)
	default:
		panic("Treat Action Problem: Server Error")
	}
	return actionResult
}

func TreatSearchAction(a *Action) *ActionResult {
	ar := &ActionResult{}

	// TODO: for korean user, use korean
	phrase := strings.Join(a.Payloads, "%20")
	glosbeParameter := &glosbe.GlosbeParameter{
		LanguageFrom: "eng",
		LanguageTo:   "jpn",
		Phrase:       phrase,
	}

	glosbeClient := glosbe.CreateGlosbeClient()
	glosbeRequest, err := glosbe.CreateGlosbeTranslateRequest(glosbeParameter)
	if err != nil {
		ar.ServerError()
		return ar
	}

	resp, err := glosbeClient.Do(glosbeRequest)
	if err != nil {
		ar.ServerError()
		return ar
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ar.ServerError()
		return ar
	}

	gRes := &glosbe.GlosbeResponse{}
	json.Unmarshal(body, gRes)

	// When phrase is not found on the glosbe site
	if len(gRes.Tucs) == 0 {
		ar.PhraseNotFound()
		return ar
	}

	mulSlice := glosbe.ExtractTenMeaning(gRes)
	ar.Text = strings.Join(mulSlice, ", ")
	ar.Status = SuccessCode
	return ar
}

func TreatSetAction(a *Action) *ActionResult {
	ar := newActionResult()

	uid := a.UserID
	timerId := a.Payloads[0]

	err := a.ValidateTime()
	if err != nil {
		fmt.Printf("Validate time error, %+v\n", err)
		ar.ServerError()
		return ar
	}

	// TODO: add quiz timer to UserInfo
	if err := timerService.AddQuizTimer(uid, timerId); err != nil {
		fmt.Printf("timerService.AddQuizTimer error, %+v\n", err)
		ar.ServerError()
		return ar
	}

	if err != userService.AddPushTimes(uid, timerId) {
		fmt.Printf("userService.AddPushTimes error, %+v\n", err)
		ar.ServerError()
		return ar
	}

	return ar
}

func TreatTimerAllAction(a *Action) *ActionResult {
	ar := newActionResult()

	uid := a.UserID
	userInfo, err := userService.ReadUserInfo(uid)
	if err != nil {
		ar.ServerError()
		return ar
	}

	ar.Text = strings.Join(userInfo.PushTimes, "\n")

	ar.Status = SuccessCode
	return ar
}

func TreatAddAction(a *Action) *ActionResult {
	ar := newActionResult()

	word, meaning := a.extractWordAndMeaning()
	uid := a.UserID
	if err := wordService.Addword(uid, word, meaning); err != nil {
		ar.ServerError()
		return ar
	}

	ar.Status = SuccessCode
	return ar
}

func TreatAllAction(a *Action) *ActionResult {
	ar := newActionResult()

	uid := a.UserID
	wordsInfo, err := wordService.ReadWords(uid)
	if err != nil {
		ar.ServerError()
		return ar
	}

	ar.Text = string(WordsAllMessage(wordsInfo))
	ar.Status = SuccessCode
	return ar
}

func TreatQuizAction(a *Action) *ActionResult {
	c := config.Setting()
	ar := newActionResult()
	ar.Text = fmt.Sprintf("%s/quiz/new/%s", c.Host, a.UserID)
	return ar
}

func TreatPredefinedAction(a *Action) *ActionResult {
	ar := newActionResult()
	bp := BotPhrase{}
	bp.Setting()
	ar.Text = string(bp[a.ActionType])
	return ar
}
