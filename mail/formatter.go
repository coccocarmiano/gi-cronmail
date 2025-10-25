package mail

import (
	"bytes"
	"cronmail/models"
	"fmt"
	"text/template"
)

type MailFormatter struct {
	mail Mail
	tmpl *template.Template
}

func GetFormatter(mail Mail) (MailFormatter, error) {
	tmpl, err := template.New("mail").Parse(mail.Body)

	if err != nil {
		return MailFormatter{}, err
	}

	return MailFormatter{
		mail: mail,
		tmpl: tmpl,
	}, nil
}

func (mf MailFormatter) String(to models.Person, params map[string]string) (string, error) {
	buf := bytes.Buffer{}
	err := mf.tmpl.Execute(&buf, params)
	if err != nil {
		return "", err
	}

	return "Content-Type: text/html\r\n" +
		fmt.Sprintf("From: %s\r\n", mf.mail.SendAs) +
		fmt.Sprintf("Reply-To: %q <%s>\r\n", mf.mail.SendAs, mf.mail.ReplyTo) +
		fmt.Sprintf("Subject: Tanti auguri, %s\r\n", to.FirstName) +
		fmt.Sprintf("To: %q <%s>\r\n", to.FirstName, to.Mail) +
		"\r\n" +
		buf.String(), nil
}
