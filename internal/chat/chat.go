package chat

type ChatEvent struct {
	Type    string  `json:"type"`
	User    User    `json:"user"`    // The user who clicked the button
	Message Message `json:"message"` // The message containing the card
	Action  `json:"action"`
}

type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type Message struct {
	Text   string `json:"text"` // Message content (if it's not just a card)
	Sender User   `json:"sender"`
	Space  struct {
		Name string `json:"name"`
	} `json:"space"`
}
type Action struct {
	ActionMethodName string            `json:"actionMethodName"` // The function name: "createGitlabIssue"
	Parameters       []ActionParameter `json:"parameters"`       // The key/value pairs
	ActionSource     string            `json:"actionSource"`     // e.g., "CARD_BUTTON"
}

type ActionParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (ce *ChatEvent) GetMessage() string {
	return ce.Message.Text
}
func (ce *ChatEvent) GetEventType() string {
	return ce.Type
}
