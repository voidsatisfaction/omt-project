package botService

import (
	"errors"
	"reflect"
	"testing"
)

func TestValidateTime(t *testing.T) {
	tests := []struct {
		expect error
		action *Action
	}{
		{
			expect: nil,
			action: createDummyAction(Set, []string{"12:34"}),
		},
		{
			expect: errors.New("error"),
			action: createDummyAction(Set, []string{"123:456"}),
		},
	}

	for _, test := range tests {
		a := test.action
		expect := test.expect
		actual := a.ValidateTime()
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("Expect: %+v, got: %+v", test.expect, actual)
		}
	}
}
