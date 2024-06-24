package actions

import (
	"fmt"

	"github.com/SineChat/notification-ms/internal/models"
	"github.com/SineChat/notification-ms/services/names"
	"github.com/gin-gonic/gin"
)

func Bind(c *gin.Context, name names.NotificationName) (interface{}, error) {
	switch name {
	case names.SendWelcomeMail:
		req := models.SendWelcomeMail{}
		err := c.ShouldBind(&req)
		return req, err
	case names.SendResetPasswordMail:
		req := models.SendResetPasswordMail{}
		err := c.ShouldBind(&req)
		return req, err
	case names.SendVerificationMail:
		req := models.SendVerificationMail{}
		err := c.ShouldBind(&req)
		return req, err
	default:
		return nil, fmt.Errorf("bind for %v, not implemented", name)
	}
}
