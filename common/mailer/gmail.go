package mailer

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type GmailConfig struct {
	SenderAddress string
	Host          string
	Port          string
	Username      string
	Password      string
}

type GmailPayload struct {
	ReceiverEmail string
	Subject       string
	Message       string
}

type Gmail interface {
	V1(payload GmailPayload)
}

type GmailCtx struct {
	payload GmailConfig
}

func NewGmail(payload GmailConfig) Gmail {
	return &GmailCtx{
		payload: payload,
	}
}

// GmailEmailSender Gmail smtp sender
func (gc *GmailCtx) V1(payload GmailPayload) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", gc.payload.SenderAddress)
	mailer.SetHeader("To", payload.ReceiverEmail)
	mailer.SetHeader("Subject", payload.Subject)
	mailer.SetBody("text/html", payload.Message)

	port, _ := strconv.Atoi(gc.payload.Port)

	dialer := gomail.NewDialer(
		gc.payload.Host,
		port,
		gc.payload.Username,
		gc.payload.Password,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		logrus.Error(err)
	}
}
