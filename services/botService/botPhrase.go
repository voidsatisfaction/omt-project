package botService

import (
	"fmt"
	"omt-project/services/wordService"
	"strings"
)

type DynamicBotMessage string
type StaticBotMessage string

type BotPhrase map[ActionType]StaticBotMessage

const (
	Invalid = "INVALID"

	InvalidCommand = "INVALID_COMMAND"
	PhraseNotFound = "PHRASE_NOT_FOUND"

	Add    = "add"
	Search = "search"
	All    = "all"
)

func (bp BotPhrase) Setting() BotPhrase {
	bp[Invalid] = `先生は、君が何を話しているかわからないの。`
	bp[InvalidCommand] = `
	こんな感じで命令してね。

	search [探したい単語] : 単語の意味を教える。
	add    [追加したい単語] [意味] : 単語集に単語を追加する。
	all : 自分が登録した単語を全部見せる。
`
	bp[PhraseNotFound] = `
		先生は、そういう単語はよくわからないわ。
	`
	return bp
}

func WordsAllMessage(wi *wordService.WordsInfo) DynamicBotMessage {
	msg := ""
	for _, w := range wi.Words {
		meaningStr := strings.Join(w.Meaning, " ")
		msg = fmt.Sprintf("%s\n%s %s", msg, w.Name, meaningStr)
	}
	return DynamicBotMessage(msg)
}
