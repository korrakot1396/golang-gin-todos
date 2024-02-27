package main

import (
	"log"
	"todo_api_gin_golang/api"
	_ "todo_api_gin_golang/cmd/docs" // Import the generated docs folder

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 	TaskTodos Service API
// @version	1.0
// @description A Task service API in Go using Gin framework
// @host localhost:8081
func main() {
	r := gin.Default()

	// Register Swagger endpoint
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
