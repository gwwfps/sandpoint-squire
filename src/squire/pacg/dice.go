package pacg

import "math/rand"

type DiceType byte

const (
	D4  DiceType = 4
	D6           = 6
	D8           = 8
	D10          = 10
	D12          = 12
)

type DiceGroup struct {
	Type     DiceType
	Number   byte
	Constant byte
}

func RollDice(groups []DiceGroup) uint16 {
	var total uint16 = 0
	for _, group := range groups {
		total += rollDiceGroup(&group)
	}
	return total
}

func rollDiceGroup(group *DiceGroup) uint16 {
	total := uint16(group.Constant)
	for i := 0; i < int(group.Number); i++ {
		total += uint16(rand.Intn(int(total)) + 1)
	}
	return total
}
