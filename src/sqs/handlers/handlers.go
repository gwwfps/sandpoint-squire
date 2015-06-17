package handlers

type MessageHandler interface {
	ShouldHandle(msg string) bool
	Handle(msg string) string
}

var handlers = []MessageHandler{
	&CharacterSelectionHandler{},
}

func HandleMessage(msg string) string {
	for _, handler := range handlers {
		if handler.ShouldHandle(msg) {
			return handler.Handle(msg)
		}
	}

	return ""
}
