package client

import (
	"chatbot-framework/internal/types"
	"log"
)

var CREATE_ISSUE_ENDPOINT = "/api/v4/create_issues/"

type GitlabClient interface {
	CreateIssue(issue types.Issue)
}

type RestGitlabClient struct {
}

func (r *RestGitlabClient) CreateIssue(issue types.Issue) {
	log.Printf("Creating issue with title: %s, projectId: %s, description: %s\n", issue.Title, issue.ProjectId,
		issue.Description)
}
