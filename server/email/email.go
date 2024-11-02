package email

import (
	"fmt"
	"log"

	env "github.com/SoftwareArch-BackstreetBoys/notification-service/configs"
	"gopkg.in/gomail.v2"
)

func SendEmail(sender string, receiver string, subject string, body string) error {
    if sender == "" {
        log.Println("Error: sender email is required")
        return fmt.Errorf("sender email is required")
    }
    if receiver == "" {
        log.Println("Error: receiver email is required")
        return fmt.Errorf("receiver email is required")
    }
    if subject == "" {
        log.Println("Error: email subject is required")
        return fmt.Errorf("email subject is required")
    }
    if body == "" {
        log.Println("Error: email body is required")
        return fmt.Errorf("email body is required")
    }

    m := gomail.NewMessage()
    m.SetHeader("From", sender)
    m.SetHeader("To", receiver)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    d := gomail.NewDialer("smtp.gmail.com", 587, sender, env.EnvAppPassword())

    if err := d.DialAndSend(m); err != nil {
        log.Printf("Failed to send email: %v", err)
        return err
    } else {
        log.Printf("Email sent to: %s", receiver)
    }
    return nil
}
