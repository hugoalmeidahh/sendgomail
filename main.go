package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	emails := []string{
		"hugoalmeidahh@gmail.com",
	}

	fmt.Print(len(emails))
	fmt.Print(emails[0])
	for i := 0; i < len(emails); i++ {
		e := []string{
			emails[i],
		}

		fmt.Print((emails[i]))
		sendEmail(e)
	}
}

func sendEmail(to []string) {
	from := "contato@hugoalmeidahh.com"

	user := "contato@hugoalmeidahh.com"
	password := "xpto@qwerty"

	addr := "mail.hugoalmeidahh.com:587"
	host := "mail.hugoalmeidahh.com"

	msg := []byte("From: from@hugoalmeidahh.com\r\n" +
		"to: no-reply@hugoalmeidahh.com.br \r\n" +
		"Subject: YOUR SUBJECT  \r\n\r\n" +
		"Hello World, Send mail with GoLang!\r\n\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}
