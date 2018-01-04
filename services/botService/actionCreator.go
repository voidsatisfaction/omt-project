package botService

import (
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Action struct {
	UserID     string
	ReplyToken string
	ActionType ActionType
	Payloads   []string
}

type ActionType string

const (
	Invalid ActionType = "INVALID"

	InvalidCommand ActionType = "INVALID_COMMAND"
	PhraseNotFound ActionType = "PHRASE_NOT_FOUND"

	Add      ActionType = "add"
	Search   ActionType = "search"
	All      ActionType = "all"
	Set      ActionType = "set"
	TimerAll ActionType = "timerall"
)

type CommandTypeMap map[string]bool

func CreateAction(uid string, msg *linebot.TextMessage, rToken string, eSrc *linebot.EventSource) *Action {
	a := &Action{}
	a.ReplyToken = rToken
	a.UserID = uid

	// If there is no msg
	if msg == nil {
		a.ActionType = Invalid
		return a
	}

	// Make a command map
	commandTypeMap := CommandTypeMap{
		"add":      true,
		"search":   true,
		"all":      true,
		"set":      true,
		"timerall": true,
	}

	// Check existance of the command
	msgSlice := strings.Split(msg.Text, " ")
	command := strings.ToLower(msgSlice[0])
	_, existCommand := commandTypeMap[command]
	if !existCommand {
		a.ActionType = InvalidCommand
		return a
	}

	a.ActionType = ActionType(command)

	a.Payloads = msgSlice[1:]
	return a
}
