package botService

import (
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Action struct {
	replyToken string
	actionType string
	payloads   []string
}

type CommandTypeMap map[string]bool

const (
	Invalid = "INVALID"

	InvalidCommand = "INVALID_COMMAND"
	Add            = "add"
	Search         = "search"
)

func CreateAction(msg *linebot.TextMessage, rToken string, eSrc *linebot.EventSource) *Action {
	a := &Action{}
	a.replyToken = rToken

	// If there is no msg
	if msg == nil {
		a.actionType = Invalid
		return a
	}

	// Make a command map
	commandTypeMap := CommandTypeMap{
		"add":    true,
		"search": true,
	}

	// Check existance of the command
	msgSlice := strings.Split(msg.Text, " ")
	command := strings.ToLower(msgSlice[0])
	_, existCommand := commandTypeMap[command]
	if !existCommand {
		a.actionType = InvalidCommand
		return a
	}

	a.actionType = command
	return a
}
