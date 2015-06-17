package slack

type OutgoingMessage struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	ChannelId string `json:"channel"`
	Body      string `json:"text"`
}

var idCounter = 0

func NewOutgoingMessage(channelId string, body string) *OutgoingMessage {
	idCounter++
	return &OutgoingMessage{
		Id:        idCounter,
		Type:      "message",
		ChannelId: channelId,
		Body:      body,
	}
}
