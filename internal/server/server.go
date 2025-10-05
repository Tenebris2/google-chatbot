package server

import (
	"chatbot-framework/internal/controller"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type ChatBotServer struct {
	Port string
}

func (s *ChatBotServer) Run() {
	r := gin.Default()

	r.GET("/ping", controller.HandlePing)
	r.POST("/issue", controller.HandleCreateIssue)

	r.Run(":" + s.Port)
}
