package main

import (
	"encoding/csv"
	"os"
)

func loadRecipient(filePath string, ch chan Recipient) error {
	defer close(ch) // close the channel after finish iterating all the content

	f, err := os.Open(filePath) // opening csv file
	if err != nil {
		return err
	}

	defer f.Close()

	r := csv.NewReader(f) // reading csv file
	records, err := r.ReadAll()

	if err != nil {
		return err
	}

	for _, records := range records[1:] {

		ch <- Recipient{
			Name:  records[0],
			Email: records[1],
		}
		// send -> consumer -> channels
	}

	return nil
}
