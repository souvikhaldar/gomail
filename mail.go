package gomail

import (
	"errors"
	"fmt"
	"net/smtp"
)

type Config struct {
	From     string
	To       []string
	Subject  string
	Body     string
	Password string
}

func New(from string, to []string, subject string, body string, password string) (error, *Config) {
	if from == "" || to == nil {
		return errors.New("To and from field can't be empty"), &Config{}
	}
	return nil, &Config{
		From:     from,
		To:       to,
		Subject:  subject,
		Body:     body,
		Password: password,
	}
}

func (c *Config) CreateMsg() []string {
	msgString := make([]string, len(c.To))
	for i, to := range c.To {
		msg := fmt.Sprintf("To: %s \r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n", to, c.Subject, c.Body)
		msgString[i] = msg
	}
	return msgString
}

func (c *Config) SendMail() error {
	msgSlice := c.CreateMsg()
	auth := smtp.PlainAuth("", c.From, c.Password, "smtp.gmail.com")
	for _, msg := range msgSlice {
		err := smtp.SendMail("smtp.gmail.com:587", auth, "joeymartin367@gmail.com", c.To, []byte(msg))
		if err != nil {
			return err
		}
	}
	return nil
}
