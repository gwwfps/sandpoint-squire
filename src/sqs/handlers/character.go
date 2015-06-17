package handlers

type CharacterSelectionHandler struct {
}

func (c *CharacterSelectionHandler) ShouldHandle(msg string) bool {
	return true
}

func (c *CharacterSelectionHandler) Handle(msg string) string {
	return "hello"
}
