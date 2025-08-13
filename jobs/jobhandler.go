package jobs

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/nishchaydeep15/go-task-api/model"
)

func EmailSender(category string, tasks *[]model.Task) {
	fmt.Println("Background Job: Sending Email for category:", category)

	// Filter tasks by category
	filtered := []model.Task{}
	for _, task := range *tasks {
		if strings.EqualFold(task.Category, category) {
			filtered = append(filtered, task)
		}
	}

	if len(filtered) == 0 {
		log.Println("No tasks found for category:", category)
		return
	}
	body := fmt.Sprintf("Email Content for category '%s':\n\n", category)
	for i, task := range filtered {
		body += fmt.Sprintf("Task no %d:\n", i+1)
		body += fmt.Sprintf("Name       : %s\n", task.Name)
		body += fmt.Sprintf("Completed  : %v\n", task.Completed)
		body += fmt.Sprintf("Search     : %s\n", task.Search)
		body += fmt.Sprintf("Created At : %s\n", task.CreatedAt.Format(time.RFC1123))
		body += fmt.Sprintf("Updated At : %s\n", task.UpdatedAt.Format(time.RFC1123))
		body += fmt.Sprintf("Accessed At: %s\n", task.AccessesAt.Format(time.RFC1123))
		body += fmt.Sprintf("Description: %s\n", task.Description)
		body += fmt.Sprintf("Category   : %s\n", task.Category)
		body += "\n"
	}

	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	to := os.Getenv("SMTP_TO")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	if from == "" || password == "" || to == "" || smtpHost == "" || smtpPort == "" {
		log.Println("SMTP credentials not set.")
		return
	}

	subject := fmt.Sprintf("Task Summary - %s", category)
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n"+
		"MIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"utf-8\"\r\n\r\n%s",
		from, to, subject, body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Println("Error sending email:", err)
		return
	}
	fmt.Println("Email sent successfully")
}
