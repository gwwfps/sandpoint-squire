package api

type BotState struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (bs BotState) AtHandle() string {
	return AtHandle(bs.Id)
}
