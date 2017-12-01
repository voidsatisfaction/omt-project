package botService

type BotPhrase map[ActionType]string

func (bp BotPhrase) Setting() BotPhrase {
	bp[Invalid] = `先生は、君が何を話しているかわからないの。`
	bp[InvalidCommand] = `
    こんな感じで命令してね。

    search [探したい単語]
  `
	return bp
}
