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

func GetTask(name string) {
	for _, task := range tasks {
		if strings.EqualFold(task.Name, name) {
			fmt.Printf("Task found is %s\n", task.Name)
			return
		}
	}
}

func DeleteTask(name string) {
	for index, task := range tasks {
		if strings.EqualFold(task.Name, name) {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Println("Task Deleted")
		}
	}
}
func main() {

	fmt.Println("Welcome to the API")
	for {
		fmt.Println("Enter 1 for Add Task")
		fmt.Println("Enter 2 for list Task")
		fmt.Println("Enter 3 for get Task")
		fmt.Println("Enter 4 for delete Task")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter the name of the Task")
			var name string
			fmt.Scanln(&name)
			AddTask(name)
		case 2:
			ListTasks()
		case 3:
			fmt.Println("enter name of the task")
			var name string
			fmt.Scanln(&name)
			GetTask(name)
		case 4:
			fmt.Println("Enter name of task to get delted")
			var name string
			fmt.Scanln(&name)
			DeleteTask(name)

		default:
			fmt.Println("Enter a valid option")
		}
	}

}
