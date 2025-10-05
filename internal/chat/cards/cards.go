package cards

import (
	"chatbot-framework/internal/controller/actions"

	gchat "google.golang.org/api/chat/v1"
)

func BuildCreateIssueCard() *gchat.GoogleAppsCardV1Card {
	return &gchat.GoogleAppsCardV1Card{
		Header: BuildHeader("Create Issue"),
		Name:   "CreateIssue",
		Sections: []*gchat.GoogleAppsCardV1Section{
			&gchat.GoogleAppsCardV1Section{
				Collapsible:               false,
				Header:                    "Create issue form",
				UncollapsibleWidgetsCount: 0,
				Widgets:                   BuildCreateIssueForm(),
			},
		},
	}
}

func BuildHeader(title string) *gchat.GoogleAppsCardV1CardHeader {
	return &gchat.GoogleAppsCardV1CardHeader{
		Title: title,
	}
}

func BuildCreateIssueForm() []*gchat.GoogleAppsCardV1Widget {
	return []*gchat.GoogleAppsCardV1Widget{
		buildIssueTitleWidget(),
		buildIssueDescWidget(),
		buildProjectWidget(),
		buildSubmissionButton(),
	}
}

func buildIssueTitleWidget() *gchat.GoogleAppsCardV1Widget {
	return &gchat.GoogleAppsCardV1Widget{
		TextInput: &gchat.GoogleAppsCardV1TextInput{
			Name:            "IssueTitle",
			Type:            "SINGLE_LINE",
			Label:           "Issue Title",
			HintText:        "Title",
			PlaceholderText: "Title",
			Value:           "", // Optional: Pre-fill if needed
		},
		TextParagraph: nil,
	}
}

func buildIssueDescWidget() *gchat.GoogleAppsCardV1Widget {
	return &gchat.GoogleAppsCardV1Widget{
		TextInput: &gchat.GoogleAppsCardV1TextInput{
			HintText:        "Description",
			Label:           "Issue Description",
			Name:            "IssueDesc",
			PlaceholderText: "Description",
			Type:            "MULTIPLE_LINE",
			Value:           "",
		},
	}
}

func buildProjectWidget() *gchat.GoogleAppsCardV1Widget {
	return &gchat.GoogleAppsCardV1Widget{
		SelectionInput: &gchat.GoogleAppsCardV1SelectionInput{
			Items: []*gchat.GoogleAppsCardV1SelectionItem{
				{
					Text:  "MyApp Frontend (ID: 123)",
					Value: "123", // GitLab project ID as value
				},
				{
					Text:  "Backend API (ID: 456)",
					Value: "456",
				},
				{
					Text:  "Mobile App (ID: 789)",
					Value: "789",
				},
			},
			Label: "Project Selection",
			Name:  "ProjectSelection",
			Type:  "DROPDOWN",
		},
	}
}

func buildSubmissionButton() *gchat.GoogleAppsCardV1Widget {
	return &gchat.GoogleAppsCardV1Widget{
		ButtonList: &gchat.GoogleAppsCardV1ButtonList{
			Buttons: []*gchat.GoogleAppsCardV1Button{
				&gchat.GoogleAppsCardV1Button{
					Text: "Create Gitlab Issue",
					OnClick: &gchat.GoogleAppsCardV1OnClick{
						Action: &gchat.GoogleAppsCardV1Action{
							Function: actions.CreateIssueCard,
							Parameters: []*gchat.GoogleAppsCardV1ActionParameter{
								{Key: "title", Value: "$issueTitle"},
								{Key: "description", Value: "$issueDescription"},
								{Key: "projectId", Value: "$projectId"},
							},
						},
					},
				},
			},
		},
	}
}
