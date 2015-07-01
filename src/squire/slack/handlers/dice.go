package handlers

import (
	"regexp"
	"strings"
)

type DiceHandler struct {
}

var rollDicePrefix = "roll dice "
var diceExp = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

func (c *DiceHandler) ShouldHandle(msg string) bool {
	return strings.HasPrefix(strings.ToLower(msg), rollDicePrefix)
}

func (c *DiceHandler) Handle(msg string) string {
	msg = msg[len(rollDicePrefix):]

	parts := strings.Split(msg, "+")

	for _, part := range parts {
		part = strings.Trim(part, " ")
		if strings.Contains(part, "d") {

		}
	}

	return "hello"
}
