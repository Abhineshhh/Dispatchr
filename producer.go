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

	for _, record := range records[1:] {
		if len(record) < 2 {
			continue // skip invalid rows
		}

		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}
		// send -> consumer -> channels
	}

	return nil
}
