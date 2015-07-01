package handlers

import "squire/slack/api"

type FallbackHandler struct {
}

func (c *FallbackHandler) ShouldHandle(msg api.ChatMessage) bool {
	return true
}

func (c *FallbackHandler) Handle(msg api.ChatMessage) string {
	return ""
}
