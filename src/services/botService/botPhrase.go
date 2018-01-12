package botService

import (
	"fmt"
	"omt-project/src/services/wordService"
	"strings"
)

type DynamicBotMessage string
type StaticBotMessage string

type BotPhrase map[ActionType]StaticBotMessage

func (bp BotPhrase) Setting() BotPhrase {
	bp[Invalid] = `先生は、君が何を話しているかわからないの。`
	bp[InvalidCommand] = `こんな感じで命令してね。

	--- 単語に関するコマンド ---
	search [探したい単語] : 単語の意味を教える。
	add    [追加したい単語] [意味] : 単語集に単語を追加する。
	all : 自分が登録した単語を全部見せる。

	--- Quizタイマーに関するコマンド ---
	set [時間 ex) 14:20, 2:34, 2:05] : Quizのタイマーを指定する。
	timerall : 登録したタイマーをすべて表示する。
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
