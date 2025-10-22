package consumer

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"

	tp "github.com/nagulvali/mailchimp/internal/pkg/types"
)


func EmailWorker(id int, ch chan tp.Recipient, wg *sync.WaitGroup) {


	defer wg.Done()

	for recipient := range ch {

		smtpHost := "localhost"
		smtpPort := "1025"

		formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\n%s\r\n", recipient.Email, "Just testing our email compaign.")
		msg := []byte(formattedMsg)
		
		fmt.Printf("Worker %d: Sending email to %s \n", id, recipient.Email)
		err := smtp.SendMail(smtpHost + ":" + smtpPort, nil, "noreply@vali.com", []string{recipient.Email}, msg)
		if err != nil {
			log.Fatal(err)
		}

		// add some delay to avoid rate limiting
		time.Sleep(50*time.Microsecond)
		fmt.Printf("Worker %d: Email sent to %s \n", id, recipient.Email)
	}

}