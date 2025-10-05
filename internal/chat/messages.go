package chat

import googlechat "google.golang.org/api/chat/v1"

func CreateMessage(cards []*googlechat.Card) *googlechat.Message {
	return &googlechat.Message{Cards: cards}
}
