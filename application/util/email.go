package util

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) {

	msg := "Subject: " + subject + "\n" + body

	auth := smtp.PlainAuth("", os.Getenv("CLIENT_EMAIL"), os.Getenv("CLIENT_APP_PWD"), os.Getenv("SMTP_DOMAIN"))

	err := smtp.SendMail(os.Getenv("SMTP_SERVER"), auth, os.Getenv("CLIENT_EMAIL"), []string{
		os.Getenv("CLIENT_EMAIL"),
	}, []byte(msg))
	if err != nil {
		fmt.Println("Error Sending Email: ", err)
	}

}
