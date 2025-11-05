package milestone

import "time"

type MilestoneCreateDto struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	IntershipId string    `json:"intership_id"`
}

type MilestoneUpdateDto struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	Status      *string    `json:"status"`
}
