package appservices

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

// MailService Send emails
type MailService struct{}

// SendEmail sends email DUH!
func (service *MailService) SendEmail(to string, subject string, body string) {
	from := os.Getenv("supportEmail")
	strBuilder := new(strings.Builder)

	strBuilder.WriteString("To: " + to + "\n")
	strBuilder.WriteString("From: Support" + "<" + from + ">\n")
	strBuilder.WriteString("Subject: FCM Tokens Report\n")
	strBuilder.WriteString(body)
	//tgwxobaqzrqxvoaj

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, os.Getenv("supportAuth"), "smtp.gmail.com"), from, []string{to}, []byte(strBuilder.String()))
	if err != nil {
		fmt.Println(err.Error())
	}
}
