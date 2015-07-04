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
		isOwner, err := pacg.IsCharacterOwner(msg.UserId)
		if err == nil {
			if isOwner {
				return true, "You have already chosen a character.", nil
			}
		} else {
			return true, "", err
		}

		fuzzyCharId := matches[1]
		char := pacg.FindCharacter(fuzzyCharId)

		reply := ""
		if char == nil {
			reply = fmt.Sprintf("Character _%s_ not found.", fuzzyCharId)
		} else {
			ownerId, err := char.GetOwnerId()
			if err == nil {
				if ownerId == "" {
					err = char.SetOwnerId(msg.UserId)
				} else {
					reply = "Character is already bound to a player."
				}
			}
			if err == nil {
				reply = fmt.Sprintf("Character set to _%s_.", char.Description())
			}
		}

		return true, reply, err
	}

	return false, "", nil
}
