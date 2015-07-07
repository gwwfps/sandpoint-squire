package pacg

import "fmt"

type CharSkill struct {
	Name  string
	Dice  byte
	Base  string
	Bonus byte
}

func (cs CharSkill) Description() string {
	if cs.Base == "" {
		return fmt.Sprintf("%s: 1d%d", cs.Name, cs.Dice)
	} else {
		return fmt.Sprintf("%s: %s+%d", cs.Name, cs.Base, cs.Bonus)
	}
}
