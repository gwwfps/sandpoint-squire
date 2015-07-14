package slack

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"squire/common"
	"squire/slack/api"
	"squire/slack/handlers"
)

const ERROR_LIMIT = 10

var LimitReachedError = errors.New("Error limit reached, terminating connection.")
var ConnectionClosedError = errors.New("Connection closed by server.")

func Connect(token string) error {
	err := api.InitiateRtm(token)
	if err != nil {
		return err
	}

	log.Println("Acquired RTM Websocket URL:", api.ConnState.WsUrl)
	log.Printf("Bot state: %+v", api.ConnState.Bot)
	log.Println("Attempting to initiate connection...")

	return listenOnSocket(api.ConnState.WsUrl)
}

func listenOnSocket(url string) error {
	headers := http.Header{}

	conn, response, err := websocket.DefaultDialer.Dial(url, headers)

	if err != nil {
		return common.AppendError("Failed to connect to RTM Websocket with response status: "+response.Status, err)
	}

	log.Println("Connection established, listening for incoming messages.")

	errorCount := 0
	for {
		err = processMessage(conn)
		if err != nil {
			if err == ConnectionClosedError {
				return err
			}
			log.Println(common.AppendError("Error processing incoming message:", err).Error())
			errorCount++
			if errorCount > ERROR_LIMIT {
				return LimitReachedError
			}
		}
	}

	return nil
}

func processMessage(conn *websocket.Conn) error {
	messageType, data, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	log.Println("Incoming message received.")
	log.Println("Message type:", messageType)

	switch messageType {
	case websocket.TextMessage:
		log.Println("Message body:", string(data))

		message := api.IncomingMessage{}
		err = message.UnmarshalJSON(data)
		if err != nil {
			return err
		}

		switch message.Type {
		case api.MessageEvent:
			chatMessage := message.Inner.(api.ChatMessage)
			if chatMessage.Body != "" {
				log.Println("Parsed message content:", chatMessage.Body)

				response := handlers.HandleMessage(chatMessage)

				if response == "" {
					log.Println("Empty response, ignoring.")
					return nil
				}

				log.Println("Responding:", response)

				replyMessage := api.NewChatReply(chatMessage, response)
				replyData, err := json.Marshal(replyMessage)
				if err != nil {
					return common.AppendError("Error marshalling reply:", err)
				}

				log.Println("Marshalled:", string(replyData))

				err = conn.WriteMessage(websocket.TextMessage, replyData)
				if err != nil {
					return common.AppendError("Error responding to message:", err)
				}
			} else {
				log.Println("Message cannot be handled yet, ignoring.")
			}
		default:
			log.Println("Cannot handle message event type, ignoring.")
		}
	case websocket.CloseMessage:
		return ConnectionClosedError
	default:
		log.Println("Cannot handle message type, ignoring.")
	}

	return nil
}
