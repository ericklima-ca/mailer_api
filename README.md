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
	To      string `json:"to,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body    string `json:"body,omitempty"`
	From    string `json:"from"`
	Token   string `json:"token"` // temporary token!
}
```
## TODOS
- [ ] Add validator hooks for package mailer.
