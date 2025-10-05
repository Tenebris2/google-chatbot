package client

import (
	"chatbot-framework/config"
	"chatbot-framework/internal/chat"
	"github.com/gin-gonic/gin"
	googlechat "google.golang.org/api/chat/v1"
	"net/http"
)

type GoogleChatClient interface {
	Start()
}

type RestGoogleChatClient struct {
	RestConfig config.Config
	RouteGroup *gin.RouterGroup
}

func (c *RestGoogleChatClient) Start() {
	router := gin.Default()
	// example route inside the group
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong from Google Chat client"})
	})
	router.POST("/issue", func(c *gin.Context) {
		c.JSON(http.StatusOK, chat.CreateMessage([]*googlechat.Card{&googlechat.Card{
			Header: &googlechat.CardHeader{Title: "hello"},
		}}))
	})

	router.Run(":" + c.RestConfig.Port)
}
