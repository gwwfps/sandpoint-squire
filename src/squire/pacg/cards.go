package pacg

import (
	"io/ioutil"
	"log"
	"math/rand"

	"gopkg.in/yaml.v2"
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

var CardStore [][]Card

func init() {
	CardStore = make([][]Card, WEAPON+1)

	populate("allies.yaml", ALLY)
	populate("armors.yaml", ARMOR)
	populate("barriers.yaml", BARRIER)
	populate("blessings.yaml", BLESSING)
	populate("henchmen.yaml", HENCHMAN)
	populate("items.yaml", ITEM)
	populate("loots.yaml", LOOT)
	populate("monsters.yaml", MONSTER)
	populate("spells.yaml", SPELL)
	populate("villains.yaml", VILLAIN)
	populate("weapons.yaml", WEAPON)
}

func populate(fileName string, cardType CardType) {
	var cards []Card

	content, err := ioutil.ReadFile("./data/cards/" + fileName)

	if err == nil {
		err = yaml.Unmarshal(content, &cards)
	}

	if err != nil {
		log.Fatalln("Cannot read data file", fileName, err)
	}

	log.Printf("Read %d cards from %s.\n", len(cards), fileName)

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
