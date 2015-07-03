package pacg

import (
	"math/rand"

	"squire/common"
)

type Card struct {
	Name   string
	Set    string
	Copies uint8
}

type CardType byte

const (
	ALLY CardType = iota
	ARMOR
	BARRIER
	BLESSING
	HENCHMAN
	ITEM
	LOOT
	MONSTER
	SPELL
	VILLAIN
	WEAPON
)

var CardStore = make([][]Card, WEAPON+1)

func init() {
	populateCards("allies.yaml", ALLY)
	populateCards("armors.yaml", ARMOR)
	populateCards("barriers.yaml", BARRIER)
	populateCards("blessings.yaml", BLESSING)
	populateCards("henchmen.yaml", HENCHMAN)
	populateCards("items.yaml", ITEM)
	populateCards("loots.yaml", LOOT)
	populateCards("monsters.yaml", MONSTER)
	populateCards("spells.yaml", SPELL)
	populateCards("villains.yaml", VILLAIN)
	populateCards("weapons.yaml", WEAPON)
}

func populateCards(fileName string, cardType CardType) {
	var cards []Card
	common.UnmarshalYAML("cards/"+fileName, &cards)
	CardStore[cardType] = cards
}

func RandomCard() *Card {
	cards := make(map[uint]Card)
	var count uint = 0
	for _, cardsOfType := range CardStore {
		for _, card := range cardsOfType {
			for i := uint8(0); i < card.Copies; i++ {
				cards[count] = card
				count++
			}
		}
	}

	return randomCard(cards)
}

func randomCard(cards map[uint]Card) *Card {
	size := len(cards)
	if size > 0 {
		index := uint(rand.Intn(size))
		card := cards[index]
		return &card
	}
	return nil
}
