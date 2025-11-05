package enum

type StatusEnum string

const (
	StatusPending  StatusEnum = "pending"
	StatusActive   StatusEnum = "active"
	StatusApproved StatusEnum = "approved"
	StatusRejected StatusEnum = "rejected"
)
