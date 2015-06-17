package main

import (
	"log"
	"os"
	"sqs/slack"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Token is required for starting the bot.")
		os.Exit(1)
	}

	err := slack.Connect(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
}
