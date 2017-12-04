package botService

import (
	"encoding/json"
	"io/ioutil"
	"omt-project/api/glosbe"
	"strings"
)

type ActionResult struct {
	Text string
}

// TreatAction is manager of actions
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
	}
	return actionResult
}

func TreatSearchAction(a *Action) *ActionResult {
	ar := &ActionResult{}

	// TODO: for korean user, use korean
	phrase := strings.Join(a.payloads, " ")
	glosbeParameter := &glosbe.GlosbeParameter{
		LanguageFrom: "eng",
		LanguageTo:   "jpn",
		Phrase:       phrase, // TODO: change it to searchable phrase
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

	// TODO: more elegant way
	ar.Text = strings.Join(glosbe.ExtractMultipleMeaning(gRes), ", ")
	return ar
}

func TreatPredefinedAction(a *Action) *ActionResult {
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
