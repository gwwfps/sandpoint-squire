package handlers

import (
	"math/rand"
	"time"
)

type MessageHandler interface {
	ShouldHandle(msg string) bool
	Handle(msg string) string
}

var handlers = []MessageHandler{
	&CardDrawHandler{},
	&FallbackHandler{},
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func HandleMessage(msg string) string {
	for _, handler := range handlers {
		if handler.ShouldHandle(msg) {
			return handler.Handle(msg)
		}
	}

	return ""
}
