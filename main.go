package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Task struct {
	Name string `json:"name"`
}

var tasks []Task

// func AddTask(name string) {
// 	task := Task{Name: name}
// 	tasks = append(tasks, task)
// 	fmt.Println("Task Added")
// }

// func ListTasks() {
// 	if len(tasks) == 0 {
// 		fmt.Println("No Task is found")
// 	}
// 	for i, task := range tasks {
// 		fmt.Printf("Task no %d is %s\n", i+1, task.Name)
// 	}
// }

// func GetTask(name string) *Task {
// 	for _, task := range tasks {
// 		if strings.EqualFold(task.Name, name) {
// 			fmt.Printf("Task found is %s\n", task.Name)
// 			return &task
// 		}
// 	}
// 	return nil
// }

// func DeleteTask(name string) {
// 	for index, task := range tasks {
// 		if strings.EqualFold(task.Name, name) {
// 			tasks = append(tasks[:index], tasks[index+1:]...)
// 			fmt.Println("Task Deleted")
// 		}
// 	}
// }

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	tasks = append(tasks, task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Message": "Task added"})
}

func ListTask(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
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
	for index, task := range tasks {
		if strings.EqualFold(task.Name, name) {
			tasks = append(tasks[:index], tasks[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Task Deleted"})
			return
		}
	}
}
func main() {
	fmt.Println("Welcome to the API")
	http.HandleFunc("/get", GetTask)
	http.HandleFunc("/list", ListTask)
	http.HandleFunc("/post", AddTask)
	http.HandleFunc("/delete", DeleteTask)

	http.ListenAndServe(":8080", nil)
}

// func main() {
// 	var input string
// 	var arg string
// 	var extra string
// 	Menu()
// 	for {
// 		fmt.Print("Enter command\n")
// 		input = " "
// 		arg = " "
// 		extra = " "
// 		fmt.Scanln(&input, &arg, &extra)
// 		if extra != " " {
// 			fmt.Println("enter a single word")
// 		} else {
// 			switch input {
// 			case "add", "a":
// 				AddTask(arg)
// 			case "l", "list":
// 				ListTasks()
// 			case "get", "g":
// 				GetTask(arg)
// 			case "delete", "d":
// 				DeleteTask(arg)
// 			default:
// 				fmt.Println("Enter a valid option")

// 			}

// 		}
// 	}
// }

// operations supported
// (add a) task-1
// list l
// delete d task-1
// get g task-1

// both not accepted
// only one

// go run main.go
// infinit

// return pointer
