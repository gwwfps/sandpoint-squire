package main

import (
	"fmt"
	"sqs/handlers"
)

func main() {
	for {
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)
		if input != "" {
			response := handlers.HandleMessage(input)
			fmt.Println(response)
		}
	}
}
