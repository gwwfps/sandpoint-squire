package pacg

import (
	"errors"
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

func IsCharacterOwner(userId string) (bool, error) {
	rconn, err := common.RedisPool.Dial()
	if err == nil {
		result, err := rconn.Do("SISMEMBER", "owners", userId)
		if err == nil && result == int64(1) {
			return true, nil
		}
	}
	return false, err
}

func (char Character) SetOwnerId(ownerId string) error {
	rconn, err := common.RedisPool.Dial()
	if err == nil {
		_, err = rconn.Do("SET", char.ownerKey(), ownerId)
		if err == nil {
			_, err = rconn.Do("SADD", "owners", ownerId)
		}
	}
	return err
}

func (char Character) GetOwnerId() (string, error) {
	rconn, err := common.RedisPool.Dial()
	if err != nil {
		return "", err
	}

	result, err := rconn.Do("GET", char.ownerKey())
	if err != nil {
		return "", err
	}
	if ownerId, ok := result.(string); ok {
		return ownerId, nil
	} else {
		return "", errors.New(fmt.Sprintf("Unknown response to GET %s: %+v", char.ownerKey(), result))
	}
}

func (char Character) ownerKey() string {
	return "owner:" + char.Key
}

func (char Character) Description() string {
	return fmt.Sprintf("%s the %s %s %s", char.Name, char.Gender, char.Race, char.Class)
}
