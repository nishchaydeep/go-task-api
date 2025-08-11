package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/nishchaydeep15/go-task-api/model"
	"github.com/nishchaydeep15/go-task-api/storage"
)

var (
	tasks []model.Task
	mutex sync.Mutex
)

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
	mutex.Lock()
	tasks = append(tasks, task)
	storage.SaveTasks(tasks)
	mutex.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Message": "Task added"})
}

func ListTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/json")

	completed := r.URL.Query().Get("completed")
	search := strings.TrimSpace(r.URL.Query().Get("search"))
	// fmt.Println("Search query:", search)
	mutex.Lock()
	defer mutex.Unlock()

	var filteredTasks []model.Task
	for _, task := range tasks {
		// fmt.Println("Checking task:", task.Name)
		if (completed == "" || (completed == "true" && task.Completed) || (completed == "false" && !task.Completed)) &&
			(search == "" || strings.Contains(strings.ToLower(task.Search), strings.ToLower(search))) {
			filteredTasks = append(filteredTasks, task)
		}
	}
	json.NewEncoder(w).Encode(filteredTasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/json")
	name := r.URL.Query().Get("name")

	mutex.Lock()
	defer mutex.Unlock()

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
	mutex.Lock()
	defer mutex.Unlock()

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
