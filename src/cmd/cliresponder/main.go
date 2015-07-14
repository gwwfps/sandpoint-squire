package main

import (
	"bufio"
	"fmt"
	"os"

	"squire/slack/api"
	"squire/slack/handlers"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()
		if input != "" {
			response := handlers.HandleMessage(api.ChatMessage{
				Body:      input,
				ChannelId: "D1234",
				UserId:    "U4321",
			})
			fmt.Println(response)
		}
		fmt.Print("> ")
	}
}
