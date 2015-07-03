package slack

import (
	"log"
	"squire/slack/api"
	"squire/slack/handlers"
)

type MessageHandler interface {
	Handle(msg api.ChatMessage) (bool, string, error)
}

var messageHandlers = []MessageHandler{
	handlers.NewCardDrawHandler(),
	handlers.NewCharSelHandler(),
	&handlers.FallbackHandler{},
}

func HandleMessage(msg api.ChatMessage) string {
	for _, handler := range messageHandlers {
		if shouldHandle, result, err := handler.Handle(msg); shouldHandle {
			if err != nil {
				log.Println("Error handling message:", err)
				return "Internal error, consult logs."
			}
			return result
		}
	}

	return ""
}
