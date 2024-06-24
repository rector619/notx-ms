package models

type SendWelcomeMail struct {
	Email string `json:"email"  validate:"required"`
}
