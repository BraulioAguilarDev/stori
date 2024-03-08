package mail

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	Sengrid *sendgrid.Client
}

func NewEmail(key string) *Email {
	return &Email{
		Sengrid: sendgrid.NewSendClient(key),
	}
}

func (e *Email) Send(to, from, subject, body string) error {
	sgto := mail.NewEmail("", to)
	sgfrom := mail.NewEmail("Stori", from)

	message := mail.NewSingleEmail(sgfrom, subject, sgto, "", body)
	_, err := e.Sengrid.Send(message)
	if err != nil {
		return err
	}

	return nil
}
