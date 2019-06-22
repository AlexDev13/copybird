package email

import (
	"fmt"
	"log"
	"net/smtp"
)

const MODULE_NAME = "email"

type Email struct{
	Config     *Config
}

func (e *Email) GetName() string {
	return MODULE_NAME
}

func (e *Email) GetConfig() interface{} {
	return &Config{}
}

func (e *Email) InitModule(_cfg interface{}) error {
	return nil
}

func (e *Email) Run() error {
	if err := e.SendEmail(); err != nil {
		return err
	}

	return nil
}

func (e *Email) Close() error {
	return nil
}

func (e *Email) SendEmail () error{

	from := e.Config.MailerUser
	pass := e.Config.MailerPassword
	to := e.Config.MailTo
	body := "Test message"
	subject := "Test Subject"

	header := ""
	header += fmt.Sprintf("From: %s\r\n", from)
	header += fmt.Sprintf("To: %s\r\n", to)
	header += fmt.Sprint("MIME-Version: 1.0\r\n")
	header += fmt.Sprint("Content-type: text/html\r\n")
	header += fmt.Sprintf("Subject: %s\r\n", subject)
	header += "\r\n" + body + "\r\n"


	err := smtp.SendMail("smtp.yandex.ru:587",
		smtp.PlainAuth("", from, pass, "smtp.yandex.ru"),
		from, []string{to}, []byte(header))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}
