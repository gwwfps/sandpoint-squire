package api

import (
	"fmt"
	"strings"
)

type ChatReply struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	ChannelId string `json:"channel"`
	Body      string `json:"text"`
}

var idCounter = 0

func NewChatReply(received ChatMessage, body string) *ChatReply {
	idCounter++

	if !received.IsDirect() {
		sep := " "
		if strings.Contains(body, "\n") {
			sep = "\n"
		}
		body = fmt.Sprintf("%s:%s%s", AtHandle(received.UserId), sep, body)
	}

	return &ChatReply{
		Id:        idCounter,
		Type:      "message",
		ChannelId: received.ChannelId,
		Body:      body,
	}
}
