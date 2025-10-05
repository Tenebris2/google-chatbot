package controller

import (
	"chatbot-framework/internal/client"
	"chatbot-framework/internal/service"
	"chatbot-framework/internal/types"

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
