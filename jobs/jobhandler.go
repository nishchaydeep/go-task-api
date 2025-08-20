package jobs

import (
	"fmt"
	"reflect"
	"time"

	"github.com/nishchaydeep15/go-task-api/config"
	"github.com/nishchaydeep15/go-task-api/mail"
	"github.com/nishchaydeep15/go-task-api/model"
	"github.com/nishchaydeep15/go-task-api/storage"
)

func StartEmailScheduler() {
	go func() {
		for {
			freq, err := config.EmailFrequencyDuration()
			if err != nil {
				fmt.Println("Invalid email frequency in config:", err)
			}
			time.Sleep(freq)
			sendGroupedEmails(config.Conf.EmailGroupBy)
		}
	}()
}

func sendGroupedEmails(groupBy string) {
	loadedTasks, err := storage.Store.GetAll()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	grouped := groupTasksByField(loadedTasks, groupBy)
	for groupKey, tasks := range grouped {
		mail.Send(mail.EmailData{
			Field:   groupBy,
			GroupBy: groupKey,
			Tasks:   tasks,
		})
	}
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
