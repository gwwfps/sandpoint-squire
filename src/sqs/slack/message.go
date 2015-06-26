package slack

import (
	"encoding/json"
	"strings"
)

type EventType byte

const (
	MessageEvent EventType = iota
	UnknownEvent
)

type IncomingMessage struct {
	Type  EventType
	Inner interface{}
}

type typeOnlyMessage struct {
	Type string `json:"type"`
}

type ChatMessage struct {
	ChannelId string `json:"channel"`
	UserId    string `json:"user"`
	Body      string `json:"text"`
}

func (message *IncomingMessage) UnmarshalJSON(data []byte) error {
	intermediate := typeOnlyMessage{}
	err := json.Unmarshal(data, &intermediate)
	if err != nil {
		return err
	}

	switch intermediate.Type {
	case "message":
		message.Type = MessageEvent
		inner := ChatMessage{}

		err = json.Unmarshal(data, &inner)
		if err != nil {
			return err
		}

		message.Inner = inner
	default:
		message.Type = UnknownEvent
	}

	return nil
}

func (chatMsg *ChatMessage) IsDirect() bool {
	return strings.HasPrefix(chatMsg.ChannelId, "D")
}

func (chatMsg *ChatMessage) IsGroup() bool {
	return strings.HasPrefix(chatMsg.ChannelId, "G")
}

func (chatMsg *ChatMessage) RealBody() string {
	body := ""
	if chatMsg.IsDirect() {
		body = chatMsg.Body
	} else if chatMsg.IsGroup() {
		botHandle := ConnectionContext.Bot.AtHandle()
		if strings.HasPrefix(chatMsg.Body, botHandle) {
			body = chatMsg.Body[len(botHandle):]
		}
	}
	return strings.Trim(body, " :\n")
}
