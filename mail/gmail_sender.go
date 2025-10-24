package mail

import (
	"cronmail/models"
	"fmt"
	"net/smtp"
)

type GmailClient struct {
	from string
	auth smtp.Auth
}

func GetGmailClient(cfg models.Config) (MailClient, error) {
	if cfg.GmailSender == "" {
		return nil, fmt.Errorf("gmail sender undefined")
	}
	if cfg.GmailAppPassword == "" {
		return nil, fmt.Errorf("gmail app password undefined")
	}

	return &GmailClient{
		auth: smtp.PlainAuth(
			"",
			cfg.GmailSender,
			cfg.GmailAppPassword,
			"smtp.gmail.com",
		),
		from: cfg.GmailSender,
	}, nil
}

func (g GmailClient) Send(to string, content string) error {
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		g.auth,
		g.from,
		[]string{to},
		[]byte(content),
	)

	if err != nil {
		return err
	}

	return nil
}
