package emailexecutor

import (
	"bytes"
	"log"
	"text/template"

	"github.com/nagulvali/mailchimp/internal/pkg/types"
)


func EmailExecutor(data types.EmailData) string {

	tmpl, err := template.ParseFiles("Templates/email.tmpl")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	var buf bytes.Buffer
	if err :=  tmpl.Execute(&buf, data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	return buf.String()

}