// AddTASK API
// @Summary Add a new task
// @Description Add a new task with name, category, and other details
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Task object"
// @Success 201 {object} map[string]string "Task added successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /tasks [post]
//GetTask API
// @Summary Get a task by name
// @Description Retrieve a task by its name
// @Tags tasks
// @Accept json
// @Produce json
// @Param name query string true "Task name"
// @Success 200 {object} model.Task "Task found"
// @Failure 404 {object} map[string]string "Task not found"
// @Router /tasks [get]
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
// DeleteTask API
// @Summary Delete a task by name
// @Description Delete a task by its name
// @Tags tasks
// @Accept json
// @Produce json
// @Param name query string true "Task name"
// @Success 200 {object} map[string]string "Task deleted successfully"
// @Failure 404 {object} map[string]string "Task not found"
// @Router /tasks [delete]

package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

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
	if strings.TrimSpace(task.Name) == "" {
		http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
		return
	}

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now
	task.AccessesAt = now

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
	category := strings.TrimSpace(r.URL.Query().Get("category"))
	// fmt.Println("Search query:", search)
	mutex.Lock()
	defer mutex.Unlock()

	// var filteredTasks []model.Task
	// for _, task := range tasks {
	// 	// fmt.Println("Checking task:", task.Name)
	// 	if (completed == "" || (completed == "true" && task.Completed) || (completed == "false" && !task.Completed)) &&
	// 		(search == "" || strings.Contains(strings.ToLower(task.Search), strings.ToLower(search))) {
	// 		filteredTasks = append(filteredTasks, task)
	// 	}
	// }

	filtered := tasks
	if completed != "" {
		completedBool := completed == "true"
		var temp []model.Task
		for _, task := range filtered {
			if task.Completed == completedBool {
				temp = append(temp, task)
			}
		}
		filtered = temp
	}
	if search != "" {
		var temp []model.Task
		for _, task := range filtered {
			if strings.Contains(strings.ToLower(task.Search), strings.ToLower(search)) {
				temp = append(temp, task)
			}
		}
		filtered = temp
	}

	if category != "" {
		var temp []model.Task
		for _, task := range filtered {
			if strings.EqualFold(task.Category, category) {
				temp = append(temp, task)
			}
		}
		filtered = temp
	}

	json.NewEncoder(w).Encode(filtered)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Application/json")
	name := r.URL.Query().Get("name")

	mutex.Lock()
	defer mutex.Unlock()

	for _, task := range tasks {
		if strings.EqualFold(task.Name, name) {
			task.AccessesAt = time.Now() // Update accessed_at time
			task.UpdatedAt = task.AccessesAt
			storage.SaveTasks(tasks) // Save updated task list
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
