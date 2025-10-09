package email

import (
	"bytes"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var smtpConfig SMTPConfig

func Init(config SMTPConfig) {
	smtpConfig = config
}

func SendEmail(to, subject, templatePath string, data interface{}) error {
	// Parse HTML template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	// Create email
	m := gomail.NewMessage()
	m.SetHeader("From", smtpConfig.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(smtpConfig.Host, smtpConfig.Port, smtpConfig.Username, smtpConfig.Password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	log.Println("Email sent to:", to)
	return nil
}
