package botService

import (
	"testing"
)

func createDummyAction(at ActionType, pl []string) *Action {
	return &Action{
		replyToken: "123",
		actionType: at,
		payloads:   pl,
	}
}

func createInvalidDummyAction() *Action {
	return &Action{
		replyToken: "123",
		actionType: Invalid,
		payloads:   []string{},
	}
}

func createInvalidCommandDummyAction() *Action {
	return &Action{
		replyToken: "123",
		actionType: InvalidCommand,
		payloads:   []string{},
	}
}

func TestTreatSearchActionSuccess(t *testing.T) {
	var tests = []struct {
		action *Action
	}{
		{createDummyAction(Search, []string{"water"})},
		{createDummyAction(Search, []string{"take", "off"})},
		{createDummyAction(Search, []string{"in", "front", "of"})},
	}

	for _, test := range tests {
		dummyAction := test.action
		actionResult := TreatAction(dummyAction)
		actual := actionResult.Text
		if len(actual) <= 0 {
			t.Errorf("Expect string size > 0, got %d on Parameter: %+v", len(actual), dummyAction.payloads)
		}
	}
}

func TestTreatAddAction(t *testing.T) {
	var tests = []struct {
		expect ActionStatusCode
		action *Action
	}{
		{
			SuccessCode,
			createDummyAction(Add, []string{"water", "水"}),
		},
		{
			SuccessCode,
			createDummyAction(Add, []string{"wind", "風"}),
		},
	}

	for _, test := range tests {
		dummyAction := test.action
		actionResult := TreatAction(dummyAction)
		actual := actionResult.Status
		if actual != SuccessCode {
			t.Errorf("Expect Status SUCCEESS, got %s", actual)
		}
	}
}

func TestTreatPredefinedAction(t *testing.T) {
	bp := BotPhrase{}
	bp.Setting()

	var tests = []struct {
		action *Action
		expect string
	}{
		{
			createInvalidDummyAction(),
			bp[Invalid],
		},
		{
			createInvalidCommandDummyAction(),
			bp[InvalidCommand],
		},
	}

	for _, test := range tests {
		actionResult := TreatAction(test.action)
		expect := test.expect
		actual := actionResult.Text
		if expect != actual {
			t.Errorf("Expect %s, got %s", expect, actual)
		}
	}
}

func TestPhraseNotFound(t *testing.T) {
	bp := BotPhrase{}
	bp.Setting()
	var tests = []struct {
		expect string
		action *Action
	}{
		{
			bp[PhraseNotFound],
			createDummyAction(Search, []string{"alksfmlwkemflekmlwfkwefw"}),
		},
		{
			bp[PhraseNotFound],
			createDummyAction(Search, []string{"asdfasda", "off", "wkejnfkwqe"}),
		},
	}

	for _, test := range tests {
		expect := test.expect
		dummyAction := test.action
		actionResult := TreatAction(dummyAction)
		actual := actionResult.Text
		if expect != actual {
			t.Errorf(
				"Expect %s, got %s on Parameter: %+v",
				expect, actual, dummyAction.payloads,
			)
		}
	}
}
