package botService

import (
	"encoding/json"
	"io/ioutil"
	"omt-project/api/glosbe"
	"omt-project/services/wordService"
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

// TreatAction is main treater of actions
func TreatAction(a *Action) *ActionResult {
	actionResult := &ActionResult{}
	switch a.actionType {
	case Invalid:
		actionResult = TreatPredefinedAction(a)
	case InvalidCommand:
		actionResult = TreatPredefinedAction(a)
	case Search:
		// Call word search api
		actionResult = TreatSearchAction(a)
	case Add:
		actionResult = TreatAddAction(a)
		// User add own phrase
	case All:
		actionResult = TreatAllAction(a)
	}
	return actionResult
}

func TreatSearchAction(a *Action) *ActionResult {
	ar := &ActionResult{}

	// TODO: for korean user, use korean
	phrase := strings.Join(a.payloads, "%20")
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

func TreatAddAction(a *Action) *ActionResult {
	ar := &ActionResult{}

	word, meaning := a.extractWordAndMeaning()
	uid := a.userID
	// TODO: pre exist words should not be added just increase priority
	if err := wordService.Addword(uid, word, meaning); err != nil {
		ar.ServerError()
		return ar
	}

	ar.Status = SuccessCode
	return ar
}

func TreatAllAction(a *Action) *ActionResult {
	ar := &ActionResult{}

	uid := a.userID
	wordsInfo, err := wordService.ReadWords(uid)
	if err != nil {
		ar.ServerError()
		return ar
	}

	ar.Text = string(WordsAllMessage(wordsInfo))
	ar.Status = SuccessCode
	return ar
}

func TreatPredefinedAction(a *Action) *ActionResult {
	ar := &ActionResult{}
	bp := BotPhrase{}
	bp.Setting()
	ar.Text = string(bp[a.actionType])
	return ar
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
