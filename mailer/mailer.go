package mailer

import (
	"encoding/json"
	"net"
	"regexp"
	"strconv"

	"gopkg.in/mail.v2"
)

type MailerService struct {
	HostPort string
}

type message struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	From    string `json:"from"`
	Token   string `json:"token"`
}

func (ms MailerService) SendMail(jsonBody []byte) error {
	var msg message

	if err := json.Unmarshal(jsonBody, &msg); err != nil {
		return err
	}

	email, fullAddr := parseAddr(msg.From)

	m := mail.NewMessage()
	m.SetHeader("From", fullAddr)
	m.SetHeader("To", msg.To)
	m.SetHeader("Subject", msg.Subject)
	m.SetBody("text/html", msg.Body)

	host, port_str, _ := net.SplitHostPort(ms.HostPort)
	port_number, _ := strconv.Atoi(port_str)
	d := mail.NewDialer(host, port_number, email, msg.Token)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func parseAddr(str string) (email, fullAddr string) {
	regexFullAddr := regexp.MustCompile(`^([\w\s<])+@([\w])+(\.com)(.br)?(>)$`)
	fullAddr = str
	if regexFullAddr.Match([]byte(fullAddr)) {
		emailPattern := regexp.MustCompile(`\s*\b[^@\s]+@[^\s]+\b\s*`)
		email = emailPattern.FindString(str)
		return
	} else {
		email = str
		return
	}
}
