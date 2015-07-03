package handlers

import (
	"strings"

	"squire/pacg"
	"squire/slack/api"
)

type CardDrawHandler struct {
	prefix string
}

func NewCardDrawHandler() *CardDrawHandler {
	return &CardDrawHandler{prefix: "draw "}
}

func (c *CardDrawHandler) Handle(msg api.ChatMessage) (bool, string, error) {
	if !strings.HasPrefix(msg.Body, c.prefix) {
		return false, "", nil
	}

	command := msg.Body[len(c.prefix):]

	switch {
	case strings.HasPrefix(command, "card"):
		card := pacg.RandomCard()
		if card != nil {
			return true, card.Name, nil
		}
	}

	return false, "", nil
}
