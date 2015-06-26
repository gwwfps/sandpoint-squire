package slack

type BotState struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (bs BotState) AtHandle() string {
	return "<@" + bs.Id + ">"
}
