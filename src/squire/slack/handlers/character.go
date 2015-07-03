package handlers

import (
	"fmt"
	"regexp"

	"squire/pacg"
	"squire/slack/api"
)

type CharSelHandler struct {
	regex *regexp.Regexp
}

func NewCharSelHandler() *CharSelHandler {
	return &CharSelHandler{regex: regexp.MustCompile(`(?:set character|i am the) (\w+)`)}
}

func (c *CharSelHandler) Handle(msg api.ChatMessage) (bool, string, error) {
	if !msg.IsDirect() {
		return false, "", nil
	}

	matches := c.regex.FindStringSubmatch(msg.Body)
	if matches != nil {
		fuzzyCharId := matches[1]
		char := pacg.FindCharacter(fuzzyCharId)

		var reply string
		if char == nil {
			reply = fmt.Sprintf("Class %s not found.", fuzzyCharId)
		} else {
			reply = fmt.Sprintf("Class set to %s.", char.Description())
		}

		return true, reply, nil
	}

	return false, "", nil
}
