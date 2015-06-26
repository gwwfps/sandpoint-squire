package handlers

type FallbackHandler struct {
}

func (c *FallbackHandler) ShouldHandle(msg string) bool {
	return true
}

func (c *FallbackHandler) Handle(msg string) string {
	return ""
}
