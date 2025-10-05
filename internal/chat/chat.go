package chat

type ChatEvent struct {
	Type    string `json:"type"`
	Message struct {
		Text   string `json:"text"`
		Sender struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
		} `json:"sender"`
		Space struct {
			Name string `json:"name"`
		} `json:"space"`
	} `json:"message"`
}

func (ce *ChatEvent) GetMessage() string {
	return ce.Message.Text
}
func (ce *ChatEvent) GetEventType() string {
	return ce.Type
}

type ChatMessage struct {
	Text string `json:"text"`
}

func CreateMessage(text string) *ChatMessage {
	return &ChatMessage{
		Text: text,
	}
}
