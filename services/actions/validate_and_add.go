package actions

import (
	"encoding/json"
	"fmt"

	"github.com/SineChat/notification-ms/external/request"
	"github.com/SineChat/notification-ms/internal/models"
	"github.com/SineChat/notification-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/notification-ms/services/names"
	"github.com/SineChat/notification-ms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateNotificationRequest(c *gin.Context, extReq request.ExternalRequest, v *validator.Validate, name string) (interface{}, error) {
	var (
		actionName = names.NotificationName(name)
	)

	req, err := Bind(c, actionName)
	if err != nil {
		return req, err
	}

	fmt.Println(req)

	err = v.Struct(req)
	if err != nil {
		return req, fmt.Errorf("%v", utility.ValidationResponse(err, v))
	}

	vr := mongodb.ValidateRequestM{Logger: extReq.Logger, Test: extReq.Test}
	err = vr.ValidateRequest(req)
	if err != nil {
		return req, err
	}

	return req, nil
}

func AddNotificationToDB(extReq request.ExternalRequest, db *mongodb.Database, name string, data interface{}) (models.NotificationRecord, error) {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return models.NotificationRecord{}, err
	}

	notificationRecord := models.NotificationRecord{
		Name:      name,
		Data:      string(dataByte),
		Attempts:  0,
		Sent:      false,
		Abandoned: false,
	}
	err = notificationRecord.CreateNotificationRecord(db)
	if err != nil {
		return models.NotificationRecord{}, err
	}

	return notificationRecord, nil
}
