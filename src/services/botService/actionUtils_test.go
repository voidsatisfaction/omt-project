package botService

import (
	"reflect"
	"testing"
)

func TestExtractWordAndMeaning(t *testing.T) {
	tests := []struct {
		action        *Action
		expectWord    string
		expectMeaning []string
	}{
		{createDummyAction(Add, []string{"water", "水"}), "water", []string{"水"}},
		{createDummyAction(Add, []string{"turn down", "拒絶する"}), "turn down", []string{"拒絶する"}},
	}

	for _, test := range tests {
		action := test.action
		w, ms := action.extractWordAndMeaning()
		if w != test.expectWord || !reflect.DeepEqual(test.expectMeaning, ms) {
			t.Errorf("Expect word: %s, got: %s\n", test.expectWord, w)
			t.Errorf("Expect meaning: %+v, got: %+v\n", test.expectMeaning, ms)
		}
	}
}
