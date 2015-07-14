package handlers

import (
	"strings"

	"squire/pacg"
	"squire/slack/api"
)

type CardDrawHandler struct {
	prefix string
}

func TryCardDrawHandler(msg api.ChatMessage) MessageHandler {
	prefix := "draw "
	if !strings.HasPrefix(msg.Body, prefix) {
		return nil
	}

	return &CardDrawHandler{prefix: prefix}
}

func (c *CardDrawHandler) Handle(msg api.ChatMessage) (string, error) {
	command := msg.Body[len(c.prefix):]

	switch {
	case strings.HasPrefix(command, "card"):
		card := pacg.RandomCard()
		if card != nil {
			return card.Name, nil
		}
	}

	return "", nil
}
