package request

type RequestCreateDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IssueId     string `json:"issue_id"`
}

type RequestUpdateDto struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type RequestReviewDto struct {
	Status         string `json:"status"`
	CompanyComment string `json:"company_comment"`
}
