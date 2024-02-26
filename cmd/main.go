package main

import (
	"log"
	"todo_api_gin_golang/api"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	taskService, err := api.NewTaskService("./db.sqlite")
	if err != nil {
		log.Fatalf("failed to create task service: %v", err)
	}

	taskHandler := api.NewTaskHandler(taskService)

	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks", taskHandler.ListTasks)
	r.GET("/tasks/:id", taskHandler.GetTaskByID)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
	r.DELETE("/tasks", taskHandler.DeleteAllTasks)
	r.PUT("/tasks/done/:id", taskHandler.MarkTaskAsDone)

	r.Run(":8081")
}
