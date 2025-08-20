package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/nishchaydeep15/go-task-api/filter"
	"github.com/nishchaydeep15/go-task-api/model"
	"github.com/nishchaydeep15/go-task-api/storage"
)

var (
	tasks []model.Task
	mutex sync.Mutex
)

func InitializeTasks(loadedTasks []model.Task) {
	mutex.Lock()
	defer mutex.Unlock()
	tasks = loadedTasks
}

// AddTASK API
// @Summary Add a new task
// @Description Add a new task with name, category, and other details
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Task object"
// @Success 201 {object} map[string]string "Task added successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /task [post]
func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(task.Name) == "" {
		http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
		return
	}

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now
	task.AccessesAt = now

	err := storage.Store.Add(&task)
	if err != nil {
		http.Error(w, "Failed to add task: "+err.Error(), http.StatusInternalServerError)
		return
	}
	storage.Store.Update(&task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task added"})
}

// ListTask API
// @Summary List all tasks
// @Description Retrieve a list of all tasks with optional filters
// @Tags tasks
// @Accept json
// @Produce json
// @Param completed query string false "Filter by completion status (true/false)"
// @Param search query string false "Search tasks by name or description"
// @Param category query string false "Filter by task category"
// @Success 200 {array} model.Task "List of tasks"
// @Router /tasks [get]
func ListTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/json")

	completed := r.URL.Query().Get("completed")
	category := strings.TrimSpace(r.URL.Query().Get("category"))
	name := strings.TrimSpace(r.URL.Query().Get("name"))
	description := strings.TrimSpace(r.URL.Query().Get("description"))
	important := r.URL.Query().Get("important")

	mutex.Lock()
	defer mutex.Unlock()

	var filtered []filter.TaskFilter
	if category != "" {
		filtered = append(filtered, filter.FieldFilter{Field: "category", Value: category})
	}
	if completed != "" {
		filtered = append(filtered, filter.FieldFilter{Field: "completed", Value: completed})
	}
	if name != "" {
		filtered = append(filtered, filter.FieldFilter{Field: "name", Value: name})
	}
	if description != "" {
		filtered = append(filtered, filter.FieldFilter{Field: "description", Value: description})
	}
	if important != "" {
		filtered = append(filtered, filter.FieldFilter{Field: "important", Value: important})
	}
	filters := filter.ApplyFilters(tasks, filtered...)
	json.NewEncoder(w).Encode(filters)
}

// GetTask API
// @Summary Get a task by name
// @Description Retrieve a task by its name
// @Tags tasks
// @Accept json
// @Produce json
// @Param name query string true "Task name"
// @Success 200 {object} model.Task "Task found"
// @Failure 404 {object} map[string]string "Task not found"
// @Router /task [get]
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Task name is required", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for _, task := range tasks {
		if strings.EqualFold(task.Name, name) {
			task.AccessesAt = time.Now() // Update accessed_at time
			task.UpdatedAt = task.AccessesAt
			storage.Store.Update(&task) // Save updated task list
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// DeleteTask API
// @Summary Delete a task by name
// @Description Delete a task by its name
// @Tags tasks
// @Accept json
// @Produce json
// @Param name query string true "Task name"
// @Success 200 {object} map[string]string "Task deleted successfully"
// @Failure 404 {object} map[string]string "Task not found"
// @Router /task [delete]
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
			storage.Store.Delete(&task)
			w.Header().Set("Content-Type", "Application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Task Deleted"})
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
