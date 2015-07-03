package handlers

import "squire/slack/api"

type FallbackHandler struct {
}

func (c *FallbackHandler) Handle(msg api.ChatMessage) (bool, string, error) {
	return true, "", nil
}
