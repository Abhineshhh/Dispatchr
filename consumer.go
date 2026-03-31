package main

import (
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup, t *template.Template) {
	defer wg.Done()

	for recipient := range ch { // reading from channel -> continuously polling the channel

		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email, "for Testing Testing TESTINGGG!!!")
		// msg := []byte(formattedMsg)

		msg, err := executeTemplate(recipient, t)
		if err != nil {
			fmt.Printf("Worker %d error parsing template for %s\n", id, recipient.Email)
			// todo : add to dlq
			continue
		}

		fmt.Printf("Worker %d is sending email to %s \n", id, recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "abhinesh@gmail.com", []string{recipient.Email}, []byte(msg))

		if err != nil {
			log.Printf("Worker %d failed to send email to %s: %v\n", id, recipient.Email, err)
			continue
		}

		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d sent email to %s \n", id, recipient.Email)

	}

}
