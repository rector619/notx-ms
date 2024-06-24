package actions

import (
	"fmt"
	"time"

	"github.com/SineChat/notification-ms/external/request"
	"github.com/SineChat/notification-ms/internal/models"
	"github.com/SineChat/notification-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/notification-ms/services/names"
)

var (
	MaxAttempts   = 3
	RetryDuration = time.Minute * time.Duration(2)
)

func GetName(name string) names.NotificationName {
	return names.NotificationName(name)
}

func handleNotificationErr(extReq request.ExternalRequest, db *mongodb.Database, notification *models.NotificationRecord, err error) error {
	notification.Attempts += 1
	if err != nil {
		notification.Sent = false
		notification.AttemptAgain = int(time.Now().Add(RetryDuration).Unix())
		extReq.Logger.Error(fmt.Sprintf("sending %v failed, Error:%v", notification.Name, err.Error()))
	} else {
		notification.Sent = true
		extReq.Logger.Info(fmt.Sprintf("sending %v successful", notification.Name))
	}

	if notification.Attempts >= MaxAttempts {
		notification.Abandoned = true
		notification.AttemptAgain = 0
	}

	if updateErr := notification.UpdateAllFields(db); updateErr != nil {
		return updateErr
	}

	return err
}
