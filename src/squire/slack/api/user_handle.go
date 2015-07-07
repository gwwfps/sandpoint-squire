package api

import "fmt"

func AtHandle(userId string) string {
	return fmt.Sprintf("<@%s>", userId)
}
