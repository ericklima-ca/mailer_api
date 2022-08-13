package models

import "log"

type MailerMock struct{}

func (s MailerMock) SendMail(_ []byte) error {
	log.Println("email sent successfully")
	return nil
}
