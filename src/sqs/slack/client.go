package slack

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"sqs/common"

	"github.com/gorilla/websocket"
)

const ERROR_LIMIT = 10

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
			log.Println(common.AppendError("Error processing incoming message:", err).Error())
			errorCount++
			if errorCount > ERROR_LIMIT {
				return errors.New("Error limit reached, terminating connection.")
			}
		}
	}

	return nil
}

func processMessage(conn *websocket.Conn) error {
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	log.Println("Incoming message received.")
	log.Println("Message type:", messageType)
	log.Println("Message body:", string(p))
	// if err = conn.WriteMessage(messageType, p); err != nil {
	// 	return err
	// }

	return nil
}
