package slack

import (
	"math/rand"
	"time"

	"squire/slack/api"
	"squire/slack/handlers"
)

type MessageHandler interface {
	ShouldHandle(msg api.ChatMessage) bool
	Handle(msg api.ChatMessage) string
}

var messageHandlers = []MessageHandler{
	&handlers.CardDrawHandler{},
	&handlers.FallbackHandler{},
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func HandleMessage(msg api.ChatMessage) string {
	for _, handler := range messageHandlers {
		if handler.ShouldHandle(msg) {
			return handler.Handle(msg)
		}
	}

	return ""
}
