package controller

import (
	"chatbot-framework/internal/chat"
	"chatbot-framework/internal/chat/cards"
	"chatbot-framework/internal/client"
	"chatbot-framework/internal/controller/actions"
	"chatbot-framework/internal/service"
	"chatbot-framework/internal/types"
	"log"
	"net/http"
	"strings"

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
		handleMessage(c, event)
	case types.CardClicked:
		handleCardClicked(c, event)
	case types.AddedToSpace:
		c.JSON(http.StatusOK, gin.H{
			"message": chat.THANK_YOU_ADD_WORKSPACE,
		})
	}
}

func handleMessage(c *gin.Context, event chat.ChatEvent) {
	message := event.GetMessage()
	log.Printf("Got message: \"%s\"\n", message)
	cmdArgs := strings.Split(message, " ")
	// @Bot start .. -> get args after @Bot
	log.Println("Command args:", cmdArgs)
	if len(cmdArgs) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not enough args",
		})
		log.Println("Not enough args:", cmdArgs)
		return
	}

	cmdArgs = cmdArgs[0:]
	for i := range cmdArgs {
		cmdArgs[i] = strings.ToUpper(cmdArgs[i])
	}

	command := cmdArgs[1] // main command ex: start
	log.Println("Handling message:", command)
	switch types.CommandType[command] {
	case types.Start:
		handleStart(c)
	case types.Create:
		handleCreate(c, command)
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": chat.INCORRECT_COMMAND,
		})
	}
}

func handleCardClicked(c *gin.Context, event chat.ChatEvent) {
	switch event.Action.ActionMethodName {
	case actions.CreateIssueCard:
		handleCreateIssueCard(c, event.Parameters)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such action method"})
	}
}

func handleStart(c *gin.Context) {
}

func handleCreate(c *gin.Context, cmd string) {
	c.JSON(200, gin.H{
		"cards": cards.BuildCreateIssueCard(),
	})
}

func handleCreateIssueCard(c *gin.Context, params []chat.ActionParameter) {
	paramsMap := make(map[string]string, len(params))

	for _, p := range params {
		paramsMap[p.Key] = p.Value
	}

	title := paramsMap["title"]
	description := paramsMap["description"]
	projectId := paramsMap["projectId"]

	issue := types.Issue{
		Title:       title,
		Description: description,
		ProjectId:   projectId,
	}
	err := service.ServiceCreateIssue(issue, &client.RestGitlabClient{})

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(200, gin.H{
		"message": "created issue",
	})
}
