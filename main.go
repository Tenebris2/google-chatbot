package main

import (
	"chatbot-framework/config"
	"chatbot-framework/internal/chat"
	"context"
)

func main() {
	ctx := context.Background()

	client := chat.NewClient(ctx)
}
