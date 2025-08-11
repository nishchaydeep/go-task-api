package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nishchaydeep15/go-task-api/model"
	"github.com/nishchaydeep15/go-task-api/storage"
)

var tasks []model.Task

func init() {
	loadedTasks, err := storage.LoadTasks()
	if err == nil {
		tasks = loadedTasks
	} else {
		tasks = []model.Task{} // Initialize with an empty slice if loading fails
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/json")
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if task.Name == "" {
		http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
		return
	}
	tasks = append(tasks, task)
	storage.SaveTasks(tasks)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Message": "Task added"})
}

func ListTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/json")
	name := r.URL.Query().Get("name")
	for _, task := range tasks {
		if strings.EqualFold(task.Name, name) {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Task name is required", http.StatusBadRequest)
		return
	}
	for index, task := range tasks {
		if strings.EqualFold(strings.TrimSpace(task.Name), strings.TrimSpace(name)) {
			tasks = append(tasks[:index], tasks[index+1:]...)
			storage.SaveTasks(tasks)
			w.Header().Set("Content-Type", "Application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Task Deleted"})
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
