package botService

import (
	"encoding/json"
	"io/ioutil"
	"omt-project/api"
)

type ActionResult struct {
	Text string
}

// TreatAction is manager of actions
func TreatAction(a *Action) *ActionResult {
	actionResult := &ActionResult{}
	switch a.actionType {
	case Invalid:
		actionResult = TreatInvalidAction(a)
	case InvalidCommand:
		actionResult = TreatInvalidAction(a)
	case Search:
		// Call word search api
		actionResult = TreatSearchAction(a)
	}
	return actionResult
}

func TreatSearchAction(a *Action) *ActionResult {
	ar := &ActionResult{}

	// TODO: for korean user, use korean
	glosbeParameter := &api.GlosbeParameter{
		LanguageFrom: "eng",
		LanguageTo:   "kor",
		Phrase:       a.payloads[0],
	}

	glosbeClient := api.CreateGlosbeClient()
	glosbeRequest, err := api.CreateGlosbeTranslateRequest(glosbeParameter)
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

	gRes := api.GlosbeResponse{}
	json.Unmarshal(body, &gRes)

	// TODO: more elegant way
	ar.Text = gRes.Tucs[0].Phrase.Text
	return ar
}

func TreatInvalidAction(a *Action) *ActionResult {
	ar := &ActionResult{}
	bp := BotPhrase{}
	bp.Setting()
	ar.Text = bp[a.actionType]
	return ar
}

func TreatInvalidCommandAction(a *Action) *ActionResult {
	ar := &ActionResult{}
	bp := BotPhrase{}
	bp.Setting()
	ar.Text = bp[a.actionType]
	return ar
}

func (ar *ActionResult) ServerError() *ActionResult {
	ar.Text = "先生、今、体の調子が悪いの"
	return ar
}
