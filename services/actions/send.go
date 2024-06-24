package actions

import (
	"fmt"

	"github.com/SineChat/notification-ms/external/request"
	"github.com/SineChat/notification-ms/internal/models"
	"github.com/SineChat/notification-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/notification-ms/services/names"
	"github.com/SineChat/notification-ms/services/notifications"
)

func Send(extReq request.ExternalRequest, db *mongodb.Database, notification *models.NotificationRecord) error {
	var (
		err  error
		req  = notifications.NewNotificationObject(extReq, db, notification)
		name = GetName(notification.Name)
	)

	switch name {
	case names.SendWelcomeMail:
		err = req.SendWelcomeMail()
	case names.SendResetPasswordMail:
		err = req.SendResetPasswordMail()
	case names.SendVerificationMail:
		err = req.SendVerificationMail()
	default:
		return handleNotificationErr(extReq, db, notification, fmt.Errorf("send for %v, not implemented", notification.Name))
	}

	return handleNotificationErr(extReq, db, notification, err)
}
