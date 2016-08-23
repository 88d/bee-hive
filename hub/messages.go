package hub

type Message struct {
	Author  string `json:"author"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func NewCloseMessage() Message {
	return Message{"", "", "close"}
}

func NewPingMessage() Message {
	return Message{"", "", "ping"}
}

func NewMessage() Message {
	return Message{"", "", "msg"}
}

func NewErrorMessage(err error) Message {
	return Message{"", err.Error(), "error"}
}
