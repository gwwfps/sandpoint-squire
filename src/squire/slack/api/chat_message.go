package api

import (
	"encoding/json"
	"strings"
)

type ChatMessage struct {
	ChannelId string `json:"channel"`
	UserId    string `json:"user"`
	Body      string `json:"text"`
}

func (message *ChatMessage) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, message)
	if err != nil {
		return err
	}

	message.parseBody()
	return nil
}

func (chatMsg *ChatMessage) IsDirect() bool {
	return strings.HasPrefix(chatMsg.ChannelId, "D")
}

func (chatMsg *ChatMessage) IsGroup() bool {
	return strings.HasPrefix(chatMsg.ChannelId, "G")
}

func (chatMsg *ChatMessage) parseBody() {
	body := ""
	if chatMsg.IsDirect() {
		body = chatMsg.Body
	} else if chatMsg.IsGroup() {
		botHandle := ConnState.Bot.AtHandle()
		if strings.HasPrefix(chatMsg.Body, botHandle) {
			body = chatMsg.Body[len(botHandle):]
		}
	}
	chatMsg.Body = strings.ToLower(strings.Trim(body, " :\n"))
}
