package producer

import (
	"encoding/csv"
	"os"

	tp "github.com/nagulvali/mailchimp/internal/pkg/types"
)

func LoadRecipient(filePath string, ch chan tp.Recipient) error {
	defer close(ch)

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	// Read records by skipping header
	for _, record := range records[1:] {
		// fmt.Println(record)
		ch <- tp.Recipient{
			Name: record[0],
			Email: record[1],
		}
		// send to consumer
	}

	return nil
}