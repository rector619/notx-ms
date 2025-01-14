package send

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/SineChat/notification-ms/external/request"
	"github.com/SineChat/notification-ms/internal/config"
	"github.com/mailgun/mailgun-go/v4"
)

type EmailRequest struct {
	ExtReq         request.ExternalRequest
	To             []string `json:"to"`
	Subject        string   `json:"subject"`
	Body           string   `json:"body"`
	AttachmentName string
	Attachment     []byte
}

func NewEmailRequest(extReq request.ExternalRequest, to []string, subject, templateFileName, baseTemplateFileName string, templateData map[string]interface{}) (*EmailRequest, error) {
	body, err := ParseTemplate(extReq, templateFileName, baseTemplateFileName, templateData)
	if err != nil {
		return &EmailRequest{}, err
	}
	return &EmailRequest{
		ExtReq:  extReq,
		To:      to,
		Subject: subject,
		Body:    body, //or parsed template
	}, nil
}

func NewSimpleEmailRequest(extReq request.ExternalRequest, to []string, subject, body string) *EmailRequest {
	return &EmailRequest{
		ExtReq:  extReq,
		To:      to,
		Subject: subject,
		Body:    body, //or parsed template
	}
}

func SendEmail(extReq request.ExternalRequest, to string, subject, templateFileName, baseTemplateFileName string, data map[string]interface{}) error {
	mailRequest, err := NewEmailRequest(extReq, []string{to}, subject, templateFileName, baseTemplateFileName, data)
	if err != nil {
		return fmt.Errorf("error getting email request, %v", err)
	}

	err = mailRequest.Send()
	if err != nil {
		return fmt.Errorf("error sending email, %v", err)
	}
	return nil
}

func SendEmailWithAttachment(extReq request.ExternalRequest, to string, subject, templateFileName, baseTemplateFileName string, data map[string]interface{}, templatePath, templateBasePath, templateName string) error {
	mailRequest, err := NewEmailRequest(extReq, []string{to}, subject, templateFileName, baseTemplateFileName, data)
	if err != nil {
		return fmt.Errorf("error getting email request, %v", err)
	}

	if templateName != "" && templatePath != "" {
		err = mailRequest.ProcessPdfAttachment(templateName, templatePath, templateBasePath, data)
		if err != nil {
			return fmt.Errorf("error getting pdf attatchment, %v", err)
		}
	}

	err = mailRequest.Send()
	if err != nil {
		return fmt.Errorf("error sending email, %v", err)
	}
	return nil
}

func (e *EmailRequest) ProcessPdfAttachment(name, templatePath, templateBasePath string, data map[string]interface{}) error {
	buffer, err := GeneratePDFFromTemplate(e.ExtReq, templatePath, templateBasePath, data)
	if err != nil {
		return err
	}

	e.AttachmentName = name
	e.Attachment = buffer

	return nil
}

func (e EmailRequest) validate() error {
	if e.Subject == "" {
		return fmt.Errorf("EMAIL::validate ==> subject is required")
	}
	if e.Body == "" {
		return fmt.Errorf("EMAIL::validate ==> body is required")
	}

	if e.To == nil {
		return fmt.Errorf("receiving email is empty")
	}

	for _, v := range e.To {
		if v == "" {
			return fmt.Errorf("receiving email is empty: %s", v)
		}

		if !strings.Contains(v, "@") {
			return fmt.Errorf("receiving email is invalid: %s", v)
		}
	}

	return nil
}

func (e *EmailRequest) Send() error {

	if err := e.validate(); err != nil {
		return err
	}

	if e.ExtReq.Test {
		return nil
	}

	err := e.sendEmailViaSMTP()

	if err != nil {
		e.ExtReq.Logger.Error("error sending email: ", err.Error())
		return err
	}
	return nil
}

func (e *EmailRequest) sendEmailViaSMTP() error {
	var (
		mailConfig = config.GetConfig().Mail
	)
	mg := mailgun.NewMailgun(mailConfig.Domain, mailConfig.PrivateApiKey)

	sender := mailConfig.SenderEmail
	subject := e.Subject
	recipient := e.To

	message := mg.NewMessage(sender, subject, "", recipient...)
	body := e.Body

	message.SetHtml(body)

	if e.AttachmentName != "" {
		message.AddBufferAttachment(e.AttachmentName, e.Attachment)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
