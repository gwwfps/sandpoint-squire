package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"squire/common"
)

type connState struct {
	WsUrl string   `json:"url"`
	Bot   BotState `json:"self"`
}

var ConnState connState

func InitiateRtm(token string) error {
	response, err := http.Get("https://slack.com/api/rtm.start?token=" + token)
	if err != nil {
		return common.AppendError("Cannot reach Slack API for initiation of RTM session:", err)
	}

	defer response.Body.Close()

	cs := connState{}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return common.AppendError("Cannot read RTM initiation response:", err)
	}

	json.Unmarshal(body, &cs)

	ConnState = cs

	return nil
}
