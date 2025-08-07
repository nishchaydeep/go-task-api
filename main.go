package main

import (
	"fmt"
	"strings"
)

type Task struct {
	Name string
}

var tasks []Task

func AddTask(name string) {
	task := Task{Name: name}
	tasks = append(tasks, task)
	fmt.Println("Task Added")
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No Task is found")
	}
	for i, task := range tasks {
		fmt.Printf("Task no %d is %s\n", i+1, task.Name)
	}
}

func GetTask(name string) *Task {
	for _, task := range tasks {
		if strings.EqualFold(task.Name, name) {
			fmt.Printf("Task found is %s\n", task.Name)
			return &task
		}
	}
	return nil
}

func DeleteTask(name string) {
	for index, task := range tasks {
		if strings.EqualFold(task.Name, name) {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Println("Task Deleted")
		}
	}
}

func Menu() {
	fmt.Println("Welcome to the API")
	fmt.Println("To add task add/a <task name>")
	fmt.Println("To list a task list/l")
	fmt.Println("to get a task get/g <task name>")
	fmt.Println("to delete task delete/d <task name>")
}
func main() {
	var input string
	var arg string
	var extra string
	Menu()
	for {
		fmt.Print("Enter command\n")
		input = " "
		arg = " "
		extra = " "
		fmt.Scanln(&input, &arg, &extra)
		if extra != " " {
			fmt.Println("enter a single word")
		} else {
			switch input {
			case "add", "a":
				AddTask(arg)
			case "l", "list":
				ListTasks()
			case "get", "g":
				GetTask(arg)
			case "delete", "d":
				DeleteTask(arg)
			default:
				fmt.Println("Enter a valid option")

			}

		}
	}
}

// operations supported
// (add a) task-1
// list l
// delete d task-1
// get g task-1

// both not accepted
// only one

// go run main.go
// infinit
