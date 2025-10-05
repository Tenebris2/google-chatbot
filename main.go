package main

import (
	"chatbot-framework/client"
	"chatbot-framework/config"
	"log"
)

func main() {
	client := client.RestGoogleChatClient{
		RestConfig: config.Config{Port: "8080"},
	}

	log.Printf("Starting client")
	client.Start()
}
