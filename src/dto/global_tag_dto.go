package dto

import (
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type GlobalTagResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	Color      string `json:"color"`
	IsApproved string `json:"is_approved"`
	UsageCount string `json:"usage_count"`
}

func GlobalTagToDto(m *model.GlobalTag) GlobalTagResponse {
	var dto GlobalTagResponse
	copier.Copy(&dto, m)
	return dto
}

func GlobalTagToListDto(m []model.GlobalTag) []GlobalTagResponse {
	out := make([]GlobalTagResponse, len(m))
	for i := range m {
		out[i] = GlobalTagToDto(&m[i])
	}
	return out
}
