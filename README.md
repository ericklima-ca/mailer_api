# Mailer API
## Intro
A simple mailing service via HTTP.
It receives messages from a POST payload and send email as requested, if the message is valid. 
The message must have the format like [Message](#message-struct).
To connect to the server, set `HOST_SMTP` environment variable.
For now, it only accepts the gmail server (setted host and port to `"smtp.gmail.com:465"`).

## MESSAGE STRUCT
``` go
type message struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	From    string `json:"from"`
	Token   string `json:"token"` // a temporary token!
}
```
## TODOS
- [ ] Add validator hooks for package mailer.
