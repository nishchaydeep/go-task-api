package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nishchaydeep15/go-task-api/config"
	_ "github.com/nishchaydeep15/go-task-api/docs"
	"github.com/nishchaydeep15/go-task-api/handler"
	"github.com/nishchaydeep15/go-task-api/jobs"
	"github.com/nishchaydeep15/go-task-api/middleware"
	"github.com/nishchaydeep15/go-task-api/storage"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	config.LoadConfig()
	config.WatchConfig(func() {
		fmt.Println("Config File reloaded dynamically")
	})

	storage.InitializeStorage()
	tasks, err := storage.Store.GetAll()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}
	handler.InitializeTasks(tasks)

	if config.Conf.SendBackgroundEmail {
		jobs.StartEmailScheduler()
	} else {
		fmt.Println("Background email sending is disabled.")
	}
	port := fmt.Sprintf(":%d", config.Conf.Port)
	fmt.Println("Starting server on port", port)

	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if r.URL.Query().Get("name") != "" {
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

	http.HandleFunc("/tasks", handler.ListTask)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.ListenAndServe(port, middleware.LoggingMiddleware(http.DefaultServeMux))
}
