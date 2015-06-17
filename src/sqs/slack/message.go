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
