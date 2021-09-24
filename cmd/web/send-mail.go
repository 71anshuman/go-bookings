package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/71anshuman/go-bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMail(msg)
		}
	}()
}

func sendMail(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "smtp-relay.sendinblue.com"
	server.Port = 587
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	server.Username = "anshuman.lawania@sendinblue.com"
	server.Password = "sZCYzGf4FIV8JQMp"

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	if m.Template != "" {
		content, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%s.html", m.Template))
		if err != nil {
			app.ErrorLog.Println(err)
		}

		mailToSend := strings.Replace(string(content), "{{%content%}}", m.Content, 1)
		email.SetBody(mail.TextHTML, mailToSend)
	} else {
		email.SetBody(mail.TextHTML, m.Content)
	}

	err = email.Send(client)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent!!")

	}
}
