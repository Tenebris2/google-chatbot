package controller

import (
	"chatbot-framework/internal/chat"
	"chatbot-framework/internal/client"
	"chatbot-framework/internal/service"
	"chatbot-framework/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func HandleCreateIssue(c *gin.Context) {
	title, _ := c.Params.Get("title")
	description, _ := c.Params.Get("description")
	project, _ := c.Params.Get("project")

	service.ServiceCreateIssue(
		types.NewIssue(title, description, project),
		&client.RestGitlabClient{})

	c.JSON(200, gin.H{
		"message": "created issue",
	})
}

func HandleChat(c *gin.Context) {
	var event chat.ChatEvent

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch types.EventTypeName[event.GetEventType()] {
	case types.Message:
		HandleMessage(c, event)
	case types.CardClicked:
		HandleCardClicked(c, event)
	case types.AddedToSpace:
		c.JSON(http.StatusOK, gin.H{
			"message": chat.THANK_YOU_ADD_WORKSPACE,
		})
	}
}

func HandleMessage(c *gin.Context, event chat.ChatEvent) {
	event.GetMessage()
}

func HandleCardClicked(c *gin.Context, event chat.ChatEvent) {

}
