package pacg

import "fmt"

type CharSkill struct {
	Name  string
	Dice  DiceType
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

func (cs CharSkill) GetDiceGroup() DiceGroup {
	if cs.Base == "" {
		return DiceGroup{cs.Dice, 1, 0}
	} else {
		return DiceGroup{D4, 0, cs.Bonus}
	}
}
