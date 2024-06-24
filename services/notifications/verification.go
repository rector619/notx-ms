package notifications

import (
	"encoding/json"
	"fmt"

	"github.com/SineChat/notification-ms/internal/models"
	"github.com/SineChat/notification-ms/services/send"
)

func (n NotificationObject) SendVerificationMail() error {
	var (
		notificationData     = models.SendVerificationMail{}
		subject              = "Please Verify Your Email Address 📩"
		templateFileName     = "verification-email.html"
		baseTemplateFileName = ""
	)

	err := json.Unmarshal([]byte(n.Notification.Data), &notificationData)
	if err != nil {
		return fmt.Errorf("error decoding saved notification data, %v", err)
	}

	data, err := ConvertToMapAndAddExtraData(notificationData, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("error converting data to map, %v", err)
	}

	return send.SendEmail(n.ExtReq, notificationData.Email, subject, templateFileName, baseTemplateFileName, data)
}
