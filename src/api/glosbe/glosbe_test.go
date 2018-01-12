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

func TestExtractTenMeaningSuccess(t *testing.T) {
	expect := []string{"abc", "def", "ghi"}
	gRes := &GlosbeResponse{
		Result: "ok",
		Tucs: []Tuc{
			createTuc("abc", "jp"),
			createTuc("def", "jp"),
			createTuc("ghi", "jp"),
			createTuc("abc", "jp"),
			createTuc("abc", "jp"),
		},
	}
	actual := ExtractTenMeaning(gRes)

	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Expect %s, Got %s", expect, actual)
	}
}

// Test it only extract ten meanings
func TestExtractTenMeaningOnlyTen(t *testing.T) {
	expect := []string{
		"a", "b", "c", "d", "e",
		"f", "g", "h", "i", "j",
	}
	gRes := &GlosbeResponse{
		Result: "ok",
		Tucs: []Tuc{
			createTuc("a", "jp"), createTuc("b", "jp"),
			createTuc("c", "jp"), createTuc("d", "jp"),
			createTuc("e", "jp"), createTuc("f", "jp"),
			createTuc("g", "jp"), createTuc("h", "jp"),
			createTuc("i", "jp"), createTuc("j", "jp"),
			createTuc("k", "jp"), createTuc("l", "jp"),
			createTuc("m", "jp"), createTuc("n", "jp"),
		},
	}
	actual := ExtractTenMeaning(gRes)

	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Expect %s, Got %s", expect, actual)
	}
}
