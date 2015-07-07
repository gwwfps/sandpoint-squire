package handlers

import (
	"fmt"
	"squire/slack/api"
)

type FallbackHandler struct {
}

func (c *FallbackHandler) Handle(msg api.ChatMessage) (bool, string, error) {
	return true, fmt.Sprintf("Unknown command: %s", msg.Body), nil
}
