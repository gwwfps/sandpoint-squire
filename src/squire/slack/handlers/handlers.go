package handlers

import (
	"log"
	"squire/slack/api"
)

type MessageHandler interface {
	Handle(msg api.ChatMessage) (string, error)
}

type TryMessageHandler func(api.ChatMessage) MessageHandler

var messageHandlers = []TryMessageHandler{
	TryCardDrawHandler,
	TryCharSelHandler,
	TryCharSkillsHandler,
	NewFallbackHandler,
}

func HandleMessage(msg api.ChatMessage) string {
	for _, tryHandler := range messageHandlers {
		handler := tryHandler(msg)
		if handler != nil {
			result, err := handler.Handle(msg)
			if err != nil {
				log.Println("Error handling message:", err)
				return "Internal error, consult logs."
			}
			return result
		}
	}

	return ""
}
