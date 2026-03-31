package main

import (
	"html/template"
	"log"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	recipientChannel := make(chan Recipient) // unbuffered channel

	go func() {
		if err := loadRecipient("./emails.csv", recipientChannel); err != nil {
			log.Printf("Failed to load recipients: %v", err)
		}
	}()

	var wg sync.WaitGroup

	workerCount := 5
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg, t)
	}

	wg.Wait()
}
