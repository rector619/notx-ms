package cronjobs

import (
	"fmt"

	"github.com/SineChat/notification-ms/external/request"
	"github.com/SineChat/notification-ms/internal/models"
	"github.com/SineChat/notification-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/notification-ms/services/actions"
)

func SendNotifications(extReq request.ExternalRequest, db *mongodb.Database) {
	notificationRecord := models.NotificationRecord{}
	notificationRecords, err := notificationRecord.GetSomeUnsentNotifications(db, 200)
	if err != nil {
		extReq.Logger.Error("error getting notificatin records: ", err.Error())
		return
	}
	fmt.Println("number of records found", len(notificationRecords))

	for _, record := range notificationRecords {
		actions.Send(extReq, db, &record)
	}
}
