package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RunningPort string
	GitlabURL   string
	CredsPath   string
	BotName     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	RunningPort := os.Getenv("RUNNING_PORT")
	GitlabURL := os.Getenv("GITLAB_URL")
	CredsPath := os.Getenv("CREDENTIALS_PATH")
	BotName := os.Getenv("BOT_NAME")

	log.Println("Loading config")
	log.Println("RunningPort:", RunningPort)
	log.Println("GitlabURL:", GitlabURL)
	log.Println("CredsPath:", CredsPath)
	log.Println("BotName:", BotName)

	return &Config{
		RunningPort: RunningPort,
		GitlabURL:   GitlabURL,
		CredsPath:   CredsPath,
		BotName:     BotName,
	}
}
