package handlers

import (
	"fmt"
	"strings"

	"squire/pacg"
	"squire/slack/api"
)

type CharSkillsHandler struct {
	command string
}

func TryCharSkillsHandler(msg api.ChatMessage) MessageHandler {
	if msg.Body != "show skills" {
		return nil
	}

	return &CharSkillsHandler{}
}

func (h *CharSkillsHandler) Handle(msg api.ChatMessage) (string, error) {
	if char, err := pacg.FindOwnedCharacter(msg.UserId); err == nil {
		if char == nil {
			return "You need to select a character first.", nil
		} else {
			skills := make([]string, len(char.Skills)+1)
			skills[0] = fmt.Sprintf("Skills for %s's character _%s_:", api.AtHandle(msg.UserId), char.Description())
			for i, skill := range char.Skills {
				skills[i+1] = skill.Description()
			}
			return strings.Join(skills, "\n"), nil
		}
	} else {
		return "", err
	}
}
