package mail

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type MailSender struct {
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(fromEmailAddress string, fromEmailPassword string) *MailSender {
	return &MailSender{
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *MailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = sender.fromEmailAddress
	e.To = to
	e.Bcc = bcc
	e.Cc = cc
	e.Subject = subject
	e.Text = []byte(content)

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed  to attach file %s: %w", f, err)
		}
	}
	smtpServerAddress := fmt.Sprintf("%s:%s", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))
	smtpHost := os.Getenv("SMTP_HOST")
	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpHost)
	return e.Send(smtpServerAddress, smtpAuth)
}
