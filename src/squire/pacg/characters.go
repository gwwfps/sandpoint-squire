package pacg

import (
	"fmt"
	"squire/common"
	"strings"
)

type Character struct {
	Key    string
	Name   string
	Gender string
	Race   string
	Class  string
}

var CharacterStore = map[string]Character{}

func init() {
	var characters []Character
	common.UnmarshalYAML("characters.yaml", &characters)

	for _, char := range characters {
		CharacterStore[char.Key] = char
	}
}

func FindCharacter(fuzzyId string) *Character {
	lowerFuzzyId := strings.ToLower(fuzzyId)

	for _, char := range CharacterStore {
		criteria := []string{char.Name, char.Class}
		for _, criterion := range criteria {
			if strings.Contains(lowerFuzzyId, strings.ToLower(criterion)) {
				return &char
			}
		}
	}

	return nil
}

func (char Character) SetOwner(ownerId string) {

}

func (char Character) GetOwnerId() (string, error) {
	rconn, err := common.RedisPool.Dial()
	if err != nil {
		return "", err
	}

	ownerId, err := rconn.Do("GET", char.ownerKey())
	return ownerId.(string), err
}

func (char Character) ownerKey() string {
	return "owner:" + char.Key
}

func (char Character) Description() string {
	desc := fmt.Sprintf("%s the %s %s %s", char.Name, char.Gender, char.Race, char.Class)

	return desc
}
