package filter

import "github.com/nishchaydeep15/go-task-api/model"

type TaskFilter interface {
	Apply(tasks []model.Task) []model.Task
}
type FieldFilter struct {
	Field string
	Value string
}

func (f FieldFilter) Apply(tasks []model.Task) []model.Task {
	var filtered []model.Task
	for _, task := range tasks {
		switch f.Field {
		case "category":
			if task.Category == f.Value {
				filtered = append(filtered, task)
			}
		case "name":
			if task.Name == f.Value {
				filtered = append(filtered, task)
			}
		case "completed":
			if (f.Value == "true" && task.Completed) || (f.Value == "false" && !task.Completed) {
				filtered = append(filtered, task)
			}
		case "description":
			if task.Description == f.Value {
				filtered = append(filtered, task)
			}
		case "important":
			if (f.Value == "true" && task.Important) || (f.Value == "false" && !task.Important) {
				filtered = append(filtered, task)
			}
		default:
			continue
		}
	}
	return filtered
}

func ApplyFilters(tasks []model.Task, filters ...TaskFilter) []model.Task {
	for _, filter := range filters {
		tasks = filter.Apply(tasks)
	}
	return tasks
}
