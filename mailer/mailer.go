package mailer

import (
	"encoding/json"
	"net"
	"strconv"

	"gopkg.in/mail.v2"
)

type MailerService struct {
	HostPort string
}
type message struct {
	To      string `json:"to,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body    string `json:"body,omitempty"`
	From    string `json:"from"`
	Token   string `json:"token"`
}

func (ms MailerService) SendMail(jsonBody []byte) error {
	var msg message

	if err := json.Unmarshal(jsonBody, &msg); err != nil {
		return err
	}

	m := mail.NewMessage()

	m.SetHeader("From", msg.From)
	m.SetHeader("To", msg.To)
	m.SetHeader("Subject", msg.Subject)
	m.SetBody("text/html", msg.Body)

	host, port_str, _ := net.SplitHostPort(ms.HostPort)
	port_number, _ := strconv.Atoi(port_str)

	d := mail.NewDialer(host, port_number, msg.From, msg.Token)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
