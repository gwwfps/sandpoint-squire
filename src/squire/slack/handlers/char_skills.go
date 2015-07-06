package handlers

import (
	"fmt"
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
			return fmt.Sprintf("%+v", char.Skills), nil
		}
	} else {
		return "", err
	}
}
