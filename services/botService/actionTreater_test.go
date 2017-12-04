package botService

import (
	"testing"
)

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
	dummyAction := &Action{
		replyToken: "123",
		actionType: Search,
		payloads:   []string{"water"},
	}

	actionResult := TreatAction(dummyAction)
	actual := actionResult.Text
	if len(actual) <= 0 {
		t.Error("Expect string size > 0, got ", len(actual))
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
