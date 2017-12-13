package botService

import "fmt"

func (a *Action) extractWordAndMeaning() (string, []string) {
	word := ""
	var meaning []string
	isWord := true
	for _, w := range a.payloads {
		for _, c := range w {
			if c > '~' {
				isWord = false
			}
			break
		}
		if !isWord {
			meaning = append(meaning, w)
		} else {
			word = fmt.Sprintf("%s %s", word, w)
		}
	}
	return word, meaning
}
