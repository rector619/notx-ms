package notifications

import (
	"github.com/SineChat/notification-ms/external/request"
	"github.com/SineChat/notification-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/notification-ms/utility"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	Db        *mongodb.Database
	Validator *validator.Validate
	Logger    *utility.Logger
	ExtReq    request.ExternalRequest
}
