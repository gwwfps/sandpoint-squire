package api

import "encoding/json"

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

		if err = json.Unmarshal(data, &inner); err != nil {
			return err
		}
		inner.ParseBody()

		message.Inner = inner
	default:
		message.Type = UnknownEvent
	}

	return nil
}
