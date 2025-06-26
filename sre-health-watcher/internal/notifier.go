package internal

import (
	"bytes"
	"fmt"
	"net/http"
	"net/smtp"
)

func SendAlert(name, url, reason string, alert AlertConfig) {
	msg := fmt.Sprintf("*[%s]* DOWN: %s\nReason: %s", name, url, reason)

	if alert.SlackWebhook != "" {
		http.Post(alert.SlackWebhook, "application/json",
			bytes.NewBuffer([]byte(fmt.Sprintf(`{"text": %q}`, msg))))
	}

	if alert.Email.SMTPServer != "" {
		auth := smtp.PlainAuth("", alert.Email.Username, alert.Email.Password, alert.Email.SMTPServer)
		body := "To: " + alert.Email.To + "\r\n" +
			"Subject: ALERT: " + name + "\r\n\r\n" +
			msg
		smtp.SendMail(
			fmt.Sprintf("%s:%d", alert.Email.SMTPServer, alert.Email.Port),
			auth,
			alert.Email.From,
			[]string{alert.Email.To},
			[]byte(body),
		)
	}
}
