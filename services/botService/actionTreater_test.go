package botService

import (
	"testing"
)

func TestTreatAction(t *testing.T) {
	expect := "용"
	dummyAction := &Action{
		replyToken: "123",
		actionType: Search,
		payloads:   []string{"dragon"},
	}

	actionResult := TreatAction(dummyAction)
	actual := actionResult.Text
	if expect != actual {
		t.Error("Expect search got ", actual)
	}
}
