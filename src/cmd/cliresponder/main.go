package main

import (
	"bufio"
	"fmt"
	"os"

	"squire/slack"
	"squire/slack/api"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()
		if input != "" {
			response := slack.HandleMessage(api.ChatMessage{Body: input})
			fmt.Println(response)
		}
		fmt.Print("> ")
	}
}
