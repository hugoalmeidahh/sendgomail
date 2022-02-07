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

	sendMailWithSSL(emails)

	fmt.Print(len(emails))
	fmt.Print(emails[0])
	for i := 0; i < len(emails); i++ {
		e := []string{
			emails[i],
		}

		fmt.Print((emails[i]))
		sendMailWithoutSSL(e)

	}
}

func sendMailWithoutSSL(to []string) {
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


func sendMailWithSSL(toMail []string) {
	fmt.Println("Send mail")

	srtMail := strings.Join(toMail, ", ")

	fmt.Println(srtMail)

	from := mail.Address{"", "contato@host.com"}
	to := mail.Address{"", srtMail}
	subj := "This is the email subject test"
	body := "This is an example body.\n With two lines."

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "mail.host.com:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", "contato@host.com", "xpto@qwerty", host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}
