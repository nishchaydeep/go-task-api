package mail

import (
	"fmt"
	"net/smtp"

	"github.com/nishchaydeep15/go-task-api/config"
)

func sendSMTP(from, to, subject, body string, cfg config.SMTPConfig) error {
	auth := smtp.PlainAuth("", from, cfg.Password, cfg.Host)

	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n\r\n" +
		body

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	return smtp.SendMail(addr, auth, from, []string{to}, []byte(msg))
}
