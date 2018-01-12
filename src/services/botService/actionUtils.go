package botService

import (
	"strings"
)

func (a *Action) extractWordAndMeaning() (string, []string) {
	var meaning []string
	var word []string
	isWord := true
	for _, w := range a.Payloads {
		for _, c := range w {
			if c > '~' {
				isWord = false
			}
			break
		}
		if !isWord {
			meaning = append(meaning, w)
		} else {
			word = append(word, w)
		}
	}
	return strings.Join(word, " "), meaning
}
