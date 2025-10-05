package client

import (
	"log"
)

var CREATE_ISSUE_ENDPOINT = "/api/v4/create_issues/"

type GitlabClient interface {
	CreateIssue()
}

type RestGitlabClient struct {
}

func (r *RestGitlabClient) CreateIssue(title, projectId, description string) {
	log.Printf("Creating issue with title: %s, projectId: %s, description: %s\n", title, projectId, description)
}
