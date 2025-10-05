package chat

import gc "google.golang.org/api/chat/v1"

func CreateMessage(cards []*gc.Card) *gc.Message {
	return &gc.Message{Cards: cards}
}

func CreateCard() *gc.Card {
	return &gc.Card{}
}
