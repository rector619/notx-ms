package router

import (
	"fmt"

	"github.com/SineChat/notification-ms/external/request"
	"github.com/SineChat/notification-ms/pkg/controller/notifications"
	"github.com/SineChat/notification-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/notification-ms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Notifications(r *gin.Engine, ApiVersion string, validator *validator.Validate, db *mongodb.Database, logger *utility.Logger) *gin.Engine {
	extReq := request.ExternalRequest{Logger: logger, Test: false}
	notificationsC := notifications.Controller{Db: db, Validator: validator, Logger: logger, ExtReq: extReq}

	notificationsUrl := r.Group(fmt.Sprintf("%v", ApiVersion))
	{
		notificationsUrl.POST("/send/:name", notificationsC.SendNotification)
	}

	return r
}
