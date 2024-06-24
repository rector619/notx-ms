package names

import "github.com/SineChat/notification-ms/utility"

type NotificationName string

const (
	SendWelcomeMail       NotificationName = "send_welcome_mail"
	SendResetPasswordMail NotificationName = "send_reset_password_mail"
	SendVerificationMail  NotificationName = "send_verification_mail"
)

func GetNames(pkgImportPath string) ([]string, error) {
	// pkgImportPath example  ./services/names
	names := []string{}
	constants, err := utility.GetConstants(pkgImportPath)
	if err != nil {
		return names, err
	}

	for _, v := range constants {
		names = append(names, v)
	}

	return names, nil
}
