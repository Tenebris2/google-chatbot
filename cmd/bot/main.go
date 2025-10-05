package main

import (
	"chatbot-framework/config"
	BotServer "chatbot-framework/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	server := BotServer.ChatBotServer{Port: cfg.RunningPort}
	server.Run()
}
