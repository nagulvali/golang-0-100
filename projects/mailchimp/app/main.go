package main

import (
	"sync"

	"github.com/nagulvali/mailchimp/internal/pkg/consumer"
	"github.com/nagulvali/mailchimp/internal/pkg/producer"
	tp "github.com/nagulvali/mailchimp/internal/pkg/types"
)


func main() {

	recipientChannel := make(chan tp.Recipient)

	go producer.LoadRecipient("data/emails.csv", recipientChannel)
	
	var wg sync.WaitGroup
	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go consumer.EmailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()

}