package handlers

import (
	"fmt"
	"regexp"

	"squire/pacg"
	"squire/slack/api"
)

type CharSelHandler struct {
	fuzzyId string
}

func TryCharSelHandler(msg api.ChatMessage) MessageHandler {
	r := regexp.MustCompile(`(?:set character|i am the) (\w+)`)
	matches := r.FindStringSubmatch(msg.Body)
	if matches == nil {
		return nil
	}
	return &CharSelHandler{fuzzyId: matches[1]}
}

func (c *CharSelHandler) Handle(msg api.ChatMessage) (string, error) {
	char, err := pacg.FindOwnedCharacter(msg.UserId)
	if err == nil {
		if char != nil {
			return fmt.Sprintf("You have already chosen _%s_ as your character.", char.Description()), nil
		}
	} else {
		return "", err
	}

	char = pacg.FindCharacter(c.fuzzyId)

	if char == nil {
		return fmt.Sprintf("Character _%s_ not found.", c.fuzzyId), nil
	} else {
		if ownerId, err := char.GetOwnerId(); err == nil {
			if ownerId == "" {
				if err = char.SetOwnerId(msg.UserId); err == nil {
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
