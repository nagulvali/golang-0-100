package consumer

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"text/template"
	"time"

	tp "github.com/nagulvali/mailchimp/internal/pkg/types"
)


func EmailWorker(id int, ch chan tp.Recipient, wg *sync.WaitGroup) {


	defer wg.Done()

	for recipient := range ch {

		smtpHost := "localhost"
		smtpPort := "1025"


		// simple formatted message
		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\n%s\r\n", recipient.Email, "Just testing our email compaign.")
		// msg := []byte(formattedMsg)

		// using custom template

		tmpl, err := template.ParseFiles("Templates/email.tmpl")
		if err != nil {
			log.Fatalf("Error parsing template: %v", err)
		}

		data := tp.EmailData{
			From: 			"no-reply@example.com",
			To:					recipient.Email,
			Name:       recipient.Name,
			Subject:    "Welcome to Mailchimp Go!",
			Body:       "We're excited to have you onboard. Explore your dashboard to start sending campaigns.",
			ButtonText: "Go to Dashboard",
			ButtonLink: "https://mailchimp-go.example.com/dashboard",
		}

		var buf bytes.Buffer
		if err :=  tmpl.Execute(&buf, data); err != nil {
			log.Fatalf("Error executing template: %v", err)
		}


		fmt.Println(buf.String())
		
		fmt.Printf("Worker %d: Sending email to %s \n", id, recipient.Email)
		err = smtp.SendMail(smtpHost + ":" + smtpPort, nil, "noreply@vali.com", []string{recipient.Email}, buf.Bytes())
		if err != nil {
			// todo: handle errors in dlq to avoid blockage
			log.Fatal(err)
		}

		// add some delay to avoid rate limiting
		time.Sleep(50*time.Microsecond)
		fmt.Printf("Worker %d: Email sent to %s \n", id, recipient.Email)
	}

}