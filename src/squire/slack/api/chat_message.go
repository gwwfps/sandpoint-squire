package api

import "strings"

type ChatMessage struct {
	ChannelId string `json:"channel"`
	UserId    string `json:"user"`
	Body      string `json:"text"`
}

func (chatMsg *ChatMessage) IsDirect() bool {
	return strings.HasPrefix(chatMsg.ChannelId, "D")
}

func (chatMsg *ChatMessage) IsGroup() bool {
	return strings.HasPrefix(chatMsg.ChannelId, "G")
}

func (chatMsg *ChatMessage) ParseBody() {
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
