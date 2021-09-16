package services

import (
	"blog-api-golang/config"
	"blog-api-golang/utils"
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
)

type MailServicePayload struct {
	to      []string
	from    string
	subject string
	body    string
}

func CreateNewMail(to []string, subject string) *MailServicePayload {
	return &MailServicePayload{
		subject: subject,
		to:      to,
		from:    config.Config.MailService.Email,
	}
}

func (payload *MailServicePayload) parseHTMLTemplate(fileName string, data interface{}) error {
	htmlTemplate, err := template.ParseFiles(fileName)

	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	if err = htmlTemplate.Execute(buffer, data); err != nil {
		return err
	}

	payload.body = buffer.String()
	return nil
}

func (payload *MailServicePayload) sendMail() error {
	body := "From: The Bidu Family<" + config.Config.MailService.Email + "> \r \n" + "To: " + strings.Join(payload.to, ", ") + "\r\nSubject: " + payload.subject + "\r\n" + utils.MIME + "\r\n" + payload.body

	SMTP := fmt.Sprintf("%s:%d", config.Config.MailService.Host, config.Config.MailService.Port)
	plainAuth := smtp.PlainAuth("", config.Config.MailService.Email, config.Config.MailService.Password, config.Config.MailService.Host)

	if err := smtp.SendMail(SMTP, plainAuth, config.Config.MailService.Email, payload.to, []byte(body)); err != nil {
		return err
	}
	return nil
}

func (payload *MailServicePayload) SendMail(templateFile string, items interface{}) error {
	err := payload.parseHTMLTemplate(templateFile, items)

	if err != nil {
		return err
	}

	if err = payload.sendMail(); err != nil {
		return err
	}

	return nil
}
