// @title Go Task API
// @version 1.0
// @description A simple API for managing tasks with categories and email notifications.
// @host localhost:8070
// @BasePath /
// @schemes http
// @contact.name Nishchay Deep

package main

import (
	"fmt"
	"net/http"

	_ "github.com/nishchaydeep15/go-task-api/docs"

	"github.com/joho/godotenv"
	"github.com/nishchaydeep15/go-task-api/handler"
	"github.com/nishchaydeep15/go-task-api/jobs"
	"github.com/nishchaydeep15/go-task-api/middleware"
	"github.com/nishchaydeep15/go-task-api/storage"
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	godotenv.Load()
}

func main() {
	fmt.Println("Welcome to the API")
	storage.LoadTasks()
	_, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	} else {
		// jobs.StartEmailScheduler("Completed")
		// jobs.StartEmailScheduler("Category")
		// jobs.StartEmailScheduler("Name")
		// jobs.StartEmailScheduler("Description")
		jobs.StartEmailScheduler("Important")
	}
	fmt.Println("Server started on port 8070")
	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if name := r.URL.Query().Get("name"); name != "" {
				handler.GetTask(w, r)
			}
		case http.MethodPost:
			handler.AddTask(w, r)
		case http.MethodDelete:
			handler.DeleteTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/tasks/", handler.ListTask)

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// http.HandleFunc("/tasks/delete?name=", handler.DeleteTask)
	http.ListenAndServe(":8070", middleware.LoggingMiddleware(http.DefaultServeMux))

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
