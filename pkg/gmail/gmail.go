package gmail

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Sender interface {
	SenderEmail(to []string, subject string, text []byte, cc []string, bcc []string) error
}

type EmailSender struct {
	Name          string
	AccountEmail  string
	PasswordEmail string
}

var (
	addressEmail = "smtp.gmail.com"
	hostEmail    = "smtp.gmail.com:587"
)

func NewEmailSender(name string, accountEmail string, passwordEmail string) Sender {
	return &EmailSender{Name: name, AccountEmail: accountEmail, PasswordEmail: passwordEmail}
}

// SenderEmail implements Sender.
func (e *EmailSender) SenderEmail(to []string, subject string, text []byte, cc []string, bcc []string) error {
	a := email.NewEmail()
	a.From = e.AccountEmail
	a.To = to
	a.Cc = cc
	a.Bcc = bcc
	a.Subject = subject
	a.Text = text

	auth := smtp.PlainAuth("", e.AccountEmail, e.PasswordEmail, hostEmail)

	return a.Send(hostEmail, auth)
}
