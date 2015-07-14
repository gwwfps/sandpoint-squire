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
	Skills []CharSkill
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

func (char Character) FindSkill(name string) *CharSkill {
	for _, skill := range char.Skills {
		if skill.Name == name {
			return &skill
		}
	}
	return nil
}

func FindOwnedCharacter(userId string) (*Character, error) {
	if rconn, err := common.RedisPool.Dial(); err == nil {
		if result, err := common.RedisString(rconn.Do("GET", charKey(userId))); err == nil {
			if char, ok := CharacterStore[result]; ok {
				return &char, nil
			} else {
				return nil, nil
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (char Character) SetOwnerId(ownerId string) error {
	rconn, err := common.RedisPool.Dial()
	if err == nil {
		rconn.Send("MULTI")
		rconn.Send("SET", char.ownerKey(), ownerId)
		rconn.Send("SET", charKey(ownerId), char.Key)
		_, err = rconn.Do("EXEC")
	}
	return err
}

func (char Character) GetOwnerId() (string, error) {
	if rconn, err := common.RedisPool.Dial(); err == nil {
		return common.RedisString(rconn.Do("GET", char.ownerKey()))
	} else {
		return "", err
	}
}

func (char Character) ownerKey() string {
	return "owner:" + char.Key
}

func charKey(ownerId string) string {
	return "char:" + ownerId
}

func (char Character) Description() string {
	return fmt.Sprintf("%s the %s %s %s", char.Name, char.Gender, char.Race, char.Class)
}
