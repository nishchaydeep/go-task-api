package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"sync"

	"github.com/nishchaydeep15/go-task-api/config"
	"github.com/nishchaydeep15/go-task-api/model"
)

var (
	tmplOnce sync.Once
	tmpl     *template.Template
)

const emailTemplate = `
Email Content for {{.Field}} = '{{.GroupBy}}':

{{range .Tasks}}
Name       : {{.Name}}
Completed  : {{.Completed}}
Created At : {{.CreatedAt.Format "Mon, 06 Apr 2003 15:05:05 MST"}}
Updated At : {{.UpdatedAt.Format "Mon, 06 Apr 2003 15:05:05 MST"}}
Accessed At: {{.AccessesAt.Format "Mon, 06 Apr 2003 15:05:05 MST"}}
Description: {{.Description}}
Category   : {{.Category}}
Important  : {{.Important}}

{{end}}
`

type EmailData struct {
	Field   string
	GroupBy string
	Tasks   []model.Task
}

func Send(data EmailData) {
	fmt.Printf("Sending Email for %s = %s\n", data.Field, data.GroupBy)

	tmplOnce.Do(func() {
		var err error
		tmpl, err = template.New("email").Parse(emailTemplate)
		if err != nil {
			fmt.Printf("Error parsing template: %v", err)
			tmpl = nil
		}
	})

	if tmpl == nil {
		fmt.Println("Email template unavailable.")
		return
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		fmt.Printf("Error executing template: %v", err)
		return
	}

	cfg := config.Conf.SMTP

	if cfg.Email == "" || cfg.Password == "" || cfg.To == "" || cfg.Host == "" || cfg.Port == 0 {
		fmt.Println("SMTP credentials not set.")
		return
	}

	subject := "Task Summary " + data.Field + ": " + data.GroupBy
	if err := sendSMTP(cfg.Email, cfg.To, subject, body.String(), cfg); err != nil {
		fmt.Printf("Error sending email: %v", err)
		return
	}

	fmt.Println("Email sent successfully")
}
