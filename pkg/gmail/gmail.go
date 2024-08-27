package gmail

import (
	"net/smtp"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
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

func NewEmailSender(config config.Config) Sender {
	return &EmailSender{Name: config.NameEmail, AccountEmail: config.Account_email, PasswordEmail: config.Password_email}
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

	auth := smtp.PlainAuth("", e.AccountEmail, e.PasswordEmail, addressEmail)

	return a.Send(hostEmail, auth)
}
