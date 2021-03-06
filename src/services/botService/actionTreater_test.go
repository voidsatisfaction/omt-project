package botService

import (
	"strings"
	"testing"
)

func createDummyAction(at ActionType, pl []string) *Action {
	return &Action{
		UserID:     "abc",
		ReplyToken: "123",
		ActionType: at,
		Payloads:   pl,
	}
}

func createInvalidDummyAction() *Action {
	return &Action{
		UserID:     "abc",
		ReplyToken: "123",
		ActionType: Invalid,
		Payloads:   []string{},
	}
}

func createInvalidCommandDummyAction() *Action {
	return &Action{
		UserID:     "abc",
		ReplyToken: "123",
		ActionType: InvalidCommand,
		Payloads:   []string{},
	}
}

func TestTreatAddAction(t *testing.T) {
	var tests = []struct {
		expect ActionStatusCode
		action *Action
	}{
		{SuccessCode, createDummyAction(Add, []string{"water", "水"})},
		{SuccessCode, createDummyAction(Add, []string{"wind", "風"})},
		{SuccessCode, createDummyAction(Add, []string{"turn", "down", "拒絶する"})},
	}

	for _, test := range tests {
		dummyAction := test.action
		actionResult := TreatAction(dummyAction)
		actual := actionResult.Status
		if actual != test.expect {
			t.Errorf("Expect Status SUCCEESS, got %s", actual)
		}
	}
}

func TestTreatAllAction(t *testing.T) {
	var tests = []struct {
		expect ActionStatusCode
		action *Action
	}{
		{
			SuccessCode,
			createDummyAction(All, []string{}),
		},
	}

	for _, test := range tests {
		dummyAction := test.action
		actualResult := TreatAllAction(dummyAction)
		actual := actualResult.Status
		if actual != test.expect {
			t.Errorf("Expect Status SUCCEESS, got %s", actual)
		}
	}
}

func TestTreatSetAction(t *testing.T) {
	var tests = []struct {
		expect ActionStatusCode
		action *Action
	}{
		{SuccessCode, createDummyAction(Set, []string{"12:34"})},
		{SuccessCode, createDummyAction(Set, []string{"02:34"})},
		{SuccessCode, createDummyAction(Set, []string{"2:34"})},
		{SuccessCode, createDummyAction(Set, []string{"2:00"})},
		{FailCode, createDummyAction(Set, []string{"123:34"})},
		{FailCode, createDummyAction(Set, []string{"1234"})},
	}

	for _, test := range tests {
		dummyAction := test.action
		actualResult := TreatAction(dummyAction)
		expect := test.expect
		actual := actualResult.Status
		if actual != test.expect {
			t.Errorf("Expect Status %+v, got %s", expect, actual)
		}
	}
}

func TestTreatTimerAllAction(t *testing.T) {
	var tests = []struct {
		expect ActionStatusCode
		action *Action
	}{
		{SuccessCode, createDummyAction(TimerAll, []string{})},
	}

	for _, test := range tests {
		dummyAction := test.action
		actualResult := TreatAction(dummyAction)
		expect := test.expect
		actual := actualResult.Status
		if actual != expect {
			t.Errorf("Expect Status %+v, got %s", expect, actual)
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
		{createInvalidDummyAction(), string(bp[Invalid])},
		{createInvalidCommandDummyAction(), string(bp[InvalidCommand])},
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
			string(bp[PhraseNotFound]),
			createDummyAction(Search, []string{"alksfmlwkemflekmlwfkwefw"}),
		},
		{
			string(bp[PhraseNotFound]),
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
				expect, actual, dummyAction.Payloads,
			)
		}
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
			t.Errorf("Expect string size > 0, got %d on Parameter: %+v", len(actual), dummyAction.Payloads)
		}
	}
}

func TestQuizActionSuccess(t *testing.T) {
	var tests = []struct {
		action *Action
	}{
		{createDummyAction(Quiz, []string{})},
		{createDummyAction(Quiz, []string{"3"})},
	}

	for _, test := range tests {
		dummyAction := test.action
		actionResult := TreatAction(dummyAction)
		actual := actionResult.Text
		if !strings.HasPrefix(actual, "http") {
			t.Errorf("Expect url, got %+v", actual)
		}
	}
}
