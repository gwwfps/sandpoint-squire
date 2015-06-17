package slack

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"sqs/common"
	"sqs/handlers"

	"github.com/gorilla/websocket"
)

const ERROR_LIMIT = 10

var LimitReachedError = errors.New("Error limit reached, terminating connection.")
var ConnectionClosedError = errors.New("Connection closed by server.")

type StartState struct {
	WsUrl string `json:"url"`
}

func Connect(token string) error {
	startState, err := initiateRtm(token)
	if err != nil {
		return err
	}

	log.Println("Acquired RTM Websocket URL:", startState.WsUrl)
	log.Println("Attempting to initiate connection...")

	err = listenOnSocket(startState.WsUrl)
	if err != nil {
		return err
	}

	return nil
}

func initiateRtm(token string) (*StartState, error) {
	response, err := http.Get("https://slack.com/api/rtm.start?token=" + token)
	if err != nil {
		return nil, common.AppendError("Cannot reach Slack API for initiation of RTM session:", err)
	}

	defer response.Body.Close()

	startState := StartState{}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, common.AppendError("Cannot read RTM initiation response:", err)
	}

	json.Unmarshal(body, &startState)

	return &startState, nil
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

		message := IncomingMessage{}
		err = message.UnmarshalJSON(data)
		if err != nil {
			return err
		}

		switch message.Type {
		case MessageEvent:
			chatMessage := message.Inner.(ChatMessage)
			if chatMessage.IsDirect() {
				response := handlers.HandleMessage(chatMessage.Body)

				log.Println("Responding:", response)

				replyMessage := NewOutgoingMessage(chatMessage.ChannelId, response)
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
				log.Println("Non-direct message cannot be handled yet, ignoring.")
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
