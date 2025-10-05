package types

type Issue struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ProjectId   string `json:"projectId"`
}

func NewIssue(title, desc, pid string) Issue {
	return Issue{
		Title:       title,
		Description: desc,
		ProjectId:   pid,
	}
}
