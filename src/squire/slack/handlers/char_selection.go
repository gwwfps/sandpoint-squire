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
	matches := c.regex.FindStringSubmatch(msg.Body)
	if matches != nil {
		reply, err := c.selectChar(matches[1], msg.UserId)
		return true, reply, err
	}

	return false, "", nil
}

func (c *CharSelHandler) selectChar(fuzzyCharId string, userId string) (string, error) {
	char, err := pacg.FindOwnedCharacter(userId)
	if err == nil {
		if char != nil {
			return fmt.Sprintf("You have already chosen _%s_ as your character.", char.Description()), nil
		}
	} else {
		return "", err
	}

	char = pacg.FindCharacter(fuzzyCharId)

	if char == nil {
		return fmt.Sprintf("Character _%s_ not found.", fuzzyCharId), nil
	} else {
		if ownerId, err := char.GetOwnerId(); err == nil {
			if ownerId == "" {
				if err = char.SetOwnerId(userId); err == nil {
					return fmt.Sprintf("Character set to _%s_.", char.Description()), nil
				} else {
					return "", err
				}
			} else {
				return "Character is already bound to a player.", nil
			}
		} else {
			return "", err
		}
	}
}
