package botService

import (
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Action struct {
	userID     string
	replyToken string
	actionType ActionType
	payloads   []string
}

type ActionType string

const (
	Invalid ActionType = "INVALID"

	InvalidCommand ActionType = "INVALID_COMMAND"
	PhraseNotFound ActionType = "PHRASE_NOT_FOUND"

	Add    ActionType = "add"
	Search ActionType = "search"
	All    ActionType = "all"
	Set    ActionType = "set"
)

type CommandTypeMap map[string]bool

func CreateAction(uid string, msg *linebot.TextMessage, rToken string, eSrc *linebot.EventSource) *Action {
	a := &Action{}
	a.replyToken = rToken
	a.userID = uid

	// If there is no msg
	if msg == nil {
		a.actionType = Invalid
		return a
	}

	// Make a command map
	commandTypeMap := CommandTypeMap{
		"add":    true,
		"search": true,
		"all":    true,
		"set":    true,
	}

	// Check existance of the command
	msgSlice := strings.Split(msg.Text, " ")
	command := strings.ToLower(msgSlice[0])
	_, existCommand := commandTypeMap[command]
	if !existCommand {
		a.actionType = InvalidCommand
		return a
	}

	a.actionType = ActionType(command)

	a.payloads = msgSlice[1:]
	return a
}
