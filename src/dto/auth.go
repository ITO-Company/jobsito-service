package dto

type SignupDto struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Name            string `json:"name"`
}

type SigninDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InternSigninDto struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	InternshipId string `json:"internship_id"`
}
