package main

import (
	"fmt"
	"github.com/jwe4/bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2" // Corrected import path
	"log"
	"os"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}()
}
func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err) // Consider returning or not proceeding if connection fails
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)

	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		templateName := m.Template
		// It's good practice to use your application's logger for debug messages
		// instead of fmt.Printf directly to stdout, especially in a web app.
		// For now, let's assume app.ErrorLog or a similar app.DebugLog exists.
		// If not, fmt.Printf is okay for temporary debugging.
		log.Printf("DEBUG: Value of m.Template: '%s'\n", templateName)
		log.Printf("DEBUG: Type of m.Template: %s\n", reflect.TypeOf(templateName))

		filePath := "" // Declare filePath outside the defer so it's accessible
		defer func() {
			if r := recover(); r != nil {
				// Log the panic value and the stack trace
				log.Printf("PANIC: Recovered in sendMsg: %v\n", r)
				log.Printf("PANIC: FilePath at time of panic/recover: %s\n", filePath) // filePath from the outer scope
				log.Printf("STACK TRACE:\n%s", string(debug.Stack()))                  // <--- GET STACK TRACE HERE

				// If you want the program to continue panicking after logging:
				// panic(r)
				// Or, if you want to handle it and not crash the mail goroutine:
				// errorLog.Printf("Recovered from panic in sendMsg: %v. Stack: %s", r, string(debug.Stack()))
			}
		}()

		filePath = fmt.Sprintf("./email-templates/%s", templateName) // Assign to the outer filePath
		log.Printf("DEBUG: Calculated filePath: '%s'\n", filePath)

		log.Println("DEBUG: Attempting to read file:", filePath)
		fileData, errFileRead := os.ReadFile(filePath) // Use a different var name for this error
		if errFileRead != nil {
			log.Printf("DEBUG: os.ReadFile returned an error: %v\n", errFileRead)
			log.Printf("DEBUG: Type of os.ReadFile error: %s\n", reflect.TypeOf(errFileRead))
			if errFileRead.Error() == "can not convert value of type string to []interface{}" {
				log.Println("DEBUG: The error from os.ReadFile matches the problematic message!")
			}
			// It's crucial to handle this error. If the template can't be read,
			// you probably shouldn't proceed to send the email with a potentially empty body.
			// You might want to log to app.ErrorLog and return.
			app.ErrorLog.Printf("Failed to read email template %s: %v", filePath, errFileRead)
			return // Exit sendMsg if template read fails
		}

		log.Printf("DEBUG: Successfully read %d bytes from %s\n", len(fileData), filePath)

		// The original os.ReadFile call is now redundant because of the debug block.
		// We can use fileData directly.
		// data, err := os.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
		// if err != nil {
		// 	app.ErrorLog.Println(err)
		// }

		mailTemplate := string(fileData) // Use fileData from the debug block
		msgToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, msgToSend)
	}

	errSend := email.Send(client) // Use a different var name for this error
	if errSend != nil {
		// Use your app.ErrorLog for consistency
		app.ErrorLog.Printf("Failed to send email to %s: %v", m.To, errSend)
	} else {
		log.Println("Email sent!")
	}
}

//func sendMsg(m models.MailData) {
//	server := mail.NewSMTPClient()
//	server.Host = "localhost"
//	server.Port = 1025
//	server.KeepAlive = false
//	server.ConnectTimeout = 10 * time.Second
//	server.SendTimeout = 10 * time.Second
//
//	client, err := server.Connect()
//	if err != nil {
//		errorLog.Println(err)
//	}
//
//	email := mail.NewMSG()
//	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
//
//	if m.Template == "" {
//		email.SetBody(mail.TextHTML, m.Content)
//	} else {
//
//		// debug code start
//
//		// debug code end
//		data, err := os.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template)) // Use os.ReadFile instead of ioutil.ReadFile
//
//		if err != nil {
//			app.ErrorLog.Println(err)
//		}
//
//		mailTemplate := string(data)
//		msgToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
//		email.SetBody(mail.TextHTML, msgToSend)
//
//	}
//
//	err = email.Send(client)
//	if err != nil {
//		log.Println(err)
//	} else {
//		log.Println("Email sent!")
//	}
//}
