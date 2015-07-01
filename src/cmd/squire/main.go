package main

import (
	"log"
	"os"
	"time"

	"squire/slack"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Token is required for starting the bot.")
	}

	for {
		err := slack.Connect(os.Args[1])
		if err != nil {
			if err == slack.ConnectionClosedError {
				log.Println("Connection closed by server, retrying in 5 seconds.")
				time.Sleep(5 * time.Second)
			} else {
				log.Fatalln(err)
			}
		}
	}
}
