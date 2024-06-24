package notifications

import (
	"github.com/SineChat/notification-ms/external/request"
	"github.com/SineChat/notification-ms/internal/models"
	"github.com/SineChat/notification-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/notification-ms/utility"
)

type NotificationObject struct {
	Notification *models.NotificationRecord
	ExtReq       request.ExternalRequest
	Db           *mongodb.Database
}

func NewNotificationObject(extReq request.ExternalRequest, db *mongodb.Database, notification *models.NotificationRecord) *NotificationObject {
	return &NotificationObject{
		ExtReq:       extReq,
		Db:           db,
		Notification: notification,
	}
}

func ConvertToMapAndAddExtraData(data interface{}, newData map[string]interface{}) (map[string]interface{}, error) {
	var (
		mapData map[string]interface{}
	)

	mapData, err := utility.StructToMap(data)
	if err != nil {
		return mapData, err
	}

	for key, value := range newData {
		mapData[key] = value
	}

	return mapData, nil
}

func thisOrThatStr(this, that string) string {
	if this == "" {
		return that
	}
	return this
}
