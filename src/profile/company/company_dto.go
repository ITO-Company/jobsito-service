package company

type CompanyUpdateDto struct {
	CompanyName *string `json:"company_name"`
	Description *string `json:"description"`
	Website     *string `json:"website"`
	Phone       *string `json:"phone"`
	Address     *string `json:"address"`
	Industry    *string `json:"industry"`
	CompanySize *string `json:"company_size"`
	LogoUrl     *string `json:"logo_url"`
}
