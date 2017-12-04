package glosbe

import (
	"reflect"
	"testing"
)

func createTuc(text string, language string) Tuc {
	return Tuc{
		Phrase: Phrase{
			Text:     text,
			Language: language,
		},
	}
}

func TestExtractMultipleMeaning(t *testing.T) {
	expect := []string{"abc", "def", "ghi"}
	gRes := &GlosbeResponse{
		Result: "ok",
		Tucs: []Tuc{
			createTuc("abc", "jp"),
			createTuc("def", "jp"),
			createTuc("ghi", "jp"),
			createTuc("abc", "jp"),
		},
	}
	actual := ExtractMultipleMeaning(gRes)

	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Expect %s, Got %s", expect, actual)
	}
}
