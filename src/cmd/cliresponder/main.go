package main

import (
	"bufio"
	"fmt"
	"os"
	"sqs/handlers"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()
		if input != "" {
			response := handlers.HandleMessage(input)
			fmt.Println(response)
		}
		fmt.Print("> ")
	}
}
