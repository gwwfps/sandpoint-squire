package handlers

import (
	"fmt"
	"squire/slack/api"
)

type FallbackHandler struct {
}

func NewFallbackHandler(msg api.ChatMessage) MessageHandler {
	return &FallbackHandler{}
}

func (c *FallbackHandler) Handle(msg api.ChatMessage) (string, error) {
	return fmt.Sprintf("Unknown command: %s", msg.Body), nil
}
