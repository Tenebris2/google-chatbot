package service

import (
	"chatbot-framework/internal/client"
	"chatbot-framework/internal/types"
)

func ServiceCreateIssue(issue types.Issue, gitlabClient client.GitlabClient) error {
	gitlabClient.CreateIssue(issue)

	return nil
}
