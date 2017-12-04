package botService

type BotPhrase map[ActionType]string

const (
	Invalid = "INVALID"

	InvalidCommand = "INVALID_COMMAND"
	PhraseNotFound = "PHRASE_NOT_FOUND"
	Add            = "add"
	Search         = "search"
)

func (bp BotPhrase) Setting() BotPhrase {
	bp[Invalid] = `先生は、君が何を話しているかわからないの。`
	bp[InvalidCommand] = `
    こんな感じで命令してね。

    search [探したい単語]
  `
	bp[PhraseNotFound] = `
		先生は、そういう単語はよくわからないわ。
	`
	return bp
}
