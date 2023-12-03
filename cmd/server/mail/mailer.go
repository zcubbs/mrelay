package mail

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/zcubbs/mrelay/cmd/server/config"
	"net/smtp"
)

type Mailer interface {
	SendMail(mail Mail) error
}

type Mail struct {
	To            []string `json:"to"`
	Cc            []string `json:"cc"`
	Bcc           []string `json:"bcc"`
	Subject       string   `json:"subject"`
	Content       string   `json:"body"`
	AttachedFiles []string `json:"attachedFiles"`
}

type DefaultMailer struct {
	userName    string
	password    string
	fromName    string
	fromAddress string
	serverHost  string
	serverPort  int
}

func NewDefaultMailer(cfg config.SmtpConfig) Mailer {
	return &DefaultMailer{
		userName:    cfg.Username,
		password:    cfg.Password,
		fromName:    cfg.FromName,
		fromAddress: cfg.FromAddress,
		serverHost:  cfg.Host,
		serverPort:  cfg.Port,
	}
}

func (dm DefaultMailer) SendMail(mail Mail) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", dm.fromName, dm.fromAddress)
	e.Subject = mail.Subject
	e.HTML = []byte(mail.Content)
	e.To = mail.To
	e.Cc = mail.Cc
	e.Bcc = mail.Bcc
	for _, file := range mail.AttachedFiles {
		if _, err := e.AttachFile(file); err != nil {
			return fmt.Errorf("failed to attach file %s: %w", file, err)
		}
	}

	smtpGmailHost := fmt.Sprintf("%s:%d", dm.serverHost, dm.serverPort)
	smtpAuth := smtp.PlainAuth("", dm.userName, dm.password, dm.serverHost)

	return e.Send(smtpGmailHost, smtpAuth)
}
