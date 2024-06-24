package models

type SendVerificationMail struct {
	Email       string `json:"email" validate:"required"`
	RedirectURL string `json:"redirect_url" validate:"required"`
}
