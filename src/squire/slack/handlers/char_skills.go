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

func NewCharSkillsHandler() *CharSkillsHandler {
	return &CharSkillsHandler{command: "show skills"}
}

func (h *CharSkillsHandler) Handle(msg api.ChatMessage) (bool, string, error) {
	if msg.Body != h.command {
		return false, "", nil
	}

	reply, err := h.showSkills(msg.UserId)
	return true, reply, err
}

func (h *CharSkillsHandler) showSkills(userId string) (string, error) {
	if char, err := pacg.FindOwnedCharacter(userId); err == nil {
		if char == nil {
			return "You need to select a character first.", nil
		} else {
			skills := make([]string, len(char.Skills)+1)
			skills[0] = fmt.Sprintf("Skills for %s's character _%s_:", api.AtHandle(userId), char.Description())
			for i, skill := range char.Skills {
				skills[i+1] = skill.Description()
			}
			return strings.Join(skills, "\n"), nil
		}
	} else {
		return "", err
	}
}
