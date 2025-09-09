package smtp

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"
)

type Email struct {
	Subject      string
	Body         string
	TimeCreated  time.Time
	TimeToBeSent time.Time
}

func (e *Email) PrintEmail() {
	fmt.Printf("Subject: %s\n", e.Subject)
	fmt.Printf("Time Created: %s\n", e.TimeCreated.Format(time.RFC822Z))
	fmt.Printf("Time To Be Sent: %s\n", e.TimeToBeSent.Format(time.RFC822Z))
	fmt.Printf("--------------------------------------------------\n")
	fmt.Printf("Body:\n%s\n", e.Body)
}

func SendEmail(e Email) {

	email := os.Getenv("EMAIL")
	password := os.Getenv("GMAIL_APP_PASSWORD")

	// Gmail's SMTP server details
	smtpHost := "smtp.gmail.com"
	smtpPort := "587" // Port 587 for TLS/STARTTLS

	// The recipient and email message
	to := []string{email}
	from := email
	msg := []byte("Subject: " + e.Subject + "\r\n" +
		"\r\n" +
		e.Body)

	// Authentication with PlainAuth
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")

}
