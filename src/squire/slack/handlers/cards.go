package handlers

import (
	"strings"

	"squire/pacg"
	"squire/slack/api"
)

type CardDrawHandler struct {
}

var drawCardPrefix = "draw "

func (c *CardDrawHandler) ShouldHandle(msg api.ChatMessage) bool {
	return strings.HasPrefix(msg.Body, drawCardPrefix)
}

func (c *CardDrawHandler) Handle(msg api.ChatMessage) string {
	command := msg.Body[len(drawCardPrefix):]

	switch {
	case strings.HasPrefix(command, "card"):
		card := pacg.RandomCard()
		return card.Name
	}

	return ""
}
