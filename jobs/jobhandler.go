package jobs

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"reflect"
	"time"

	"github.com/nishchaydeep15/go-task-api/model"
	"github.com/nishchaydeep15/go-task-api/storage"
)

func StartEmailScheduler(groupBy string) {
	go func() {
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			sendGroupedEmails(groupBy)
		}
	}()
}
func groupTasksByField(tasks []model.Task, fieldName string) map[string][]model.Task {
	grouped := make(map[string][]model.Task)

	for _, task := range tasks {
		v := reflect.ValueOf(task)
		fieldVal := v.FieldByName(fieldName)

		if !fieldVal.IsValid() {
			fmt.Printf("Invalid field: %s\n", fieldName)
			continue
		}

		key := fmt.Sprintf("%v", fieldVal.Interface())
		grouped[key] = append(grouped[key], task)
	}
	return grouped
}

func sendGroupedEmails(groupBy string) {
	loadedTasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	grouped := groupTasksByField(loadedTasks, groupBy)
	for category, tasks := range grouped {
		EmailSender(category, groupBy, &tasks)
	}
}

const emailTemplate = `
Email Content for {{.Field}} = '{{.GroupBy}}':

{{range .Tasks}}
Name       : {{.Name}}
Completed  : {{.Completed}}
Created At : {{.CreatedAt.Format "Mon, 06 April 2003 15:05:05 MST"}}
Updated At : {{.UpdatedAt.Format "Mon, 06 April 2003 15:05:05 MST"}}
Accessed At: {{.AccessesAt.Format "Mon, 06 April 2003 15:05:05 MST"}}
Description: {{.Description}}
Category   : {{.Category}}
Important  : {{.Important}}

{{end}}
`

func EmailSender(groupBy string, field string, tasks *[]model.Task) {
	fmt.Printf("Sending Email for %s = %s\n", field, groupBy)
	type emailData struct {
		Field   string
		GroupBy string
		Tasks   []model.Task
	}

	if len(*tasks) == 0 {
		fmt.Printf("No tasks found for %s = %s\n", field, groupBy)
		return
	}

	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		fmt.Println("Error parsing email template:", err)
		return
	}

	data := emailData{
		Field:   field,
		GroupBy: groupBy,
		Tasks:   *tasks,
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		fmt.Println("Error executing email template:", err)
		return
	}

	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	to := os.Getenv("SMTP_TO")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	if from == "" || password == "" || to == "" || smtpHost == "" || smtpPort == "" {
		fmt.Println("SMTP credentials not set.")
		return
	}

	subject := "Task Summary - " + field + ": " + groupBy
	message := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n\r\n" +
		body.String()

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully")
}
