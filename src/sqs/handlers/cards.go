package handlers

import (
	"sqs/pacg"
	"strings"
)

type CardDrawHandler struct {
}

var drawCardPrefix = "draw "

func (c *CardDrawHandler) ShouldHandle(msg string) bool {
	return strings.HasPrefix(msg, drawCardPrefix)
}

func (c *CardDrawHandler) Handle(msg string) string {
	command := msg[len(drawCardPrefix):]

	switch {
	case strings.HasPrefix(command, "card"):
		card := pacg.RandomCard()
		return card.Name
	}

	return ""
}
