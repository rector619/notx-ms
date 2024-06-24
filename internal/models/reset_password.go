package models

type SendResetPasswordMail struct {
	Email       string `json:"email"  validate:"required"`
	RedirectURL string `json:"redirect_url" validate:"required"`
}
