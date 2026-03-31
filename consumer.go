package main

import (
	"fmt"
	"sync"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipient := range ch { // reading from channel -> continuously polling the channel
		fmt.Println(id, recipient)
	}

}
