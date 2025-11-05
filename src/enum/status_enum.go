package enum

type StatusEnum string

const (
	StatusPending  StatusEnum = "pending"
	StatusApproved StatusEnum = "approved"
	StatusRejected StatusEnum = "rejected"
)
