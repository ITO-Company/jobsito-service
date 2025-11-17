package issue

import "time"

type IssueCreateDto struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	MilestoneId string    `json:"milestone_id"`
}

type IssueUpdateDto struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	Status      *string    `json:"status"`
}
