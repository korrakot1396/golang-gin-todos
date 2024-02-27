package api

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"todo_api_gin_golang/models"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

type TaskHandlerInterface interface {
	CreateTask(c *gin.Context)
	ListTasks(c *gin.Context)
	GetTaskByID(c *gin.Context)
	UpdateTask(c *gin.Context)
	MarkTaskAsDone(c *gin.Context)
	DeleteTask(c *gin.Context)
	DeleteAllTasks(c *gin.Context)
}

type TaskHandler struct {
	Service TaskServiceInterface
}

const MaxFileSize = 10 * 1024 * 1024 // 10 MB

func NewTaskHandler(service TaskServiceInterface) *TaskHandler {
	return &TaskHandler{Service: service}
}

// CreateTask handles the creation of a new task.
// @Summary Create a new task
// @Description Create a new task in the system
// @Accept json
// @Produce json
// @Param task body models.Task true "Task object to be created"
// @Success 201 {object} models.Task "Created task"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if utf8.RuneCountInString(task.Title) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title must not exceed 100 characters"})
		return
	}

	if task.Image != "" {
		file, err := c.FormFile("image")
		if err != nil {
			if err != http.ErrMissingFile {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process image file"})
				return
			}
		} else {
			imageBase64, err := extractBase64FromFile(file)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image file format"})
				return
			}
			task.Image = imageBase64
		}
	}

	createdTask, err := h.Service.CreateTaskQuery(c.Request.Context(), &task)
	if err != nil {
		if strings.Contains(err.Error(), "task with the same title already exists") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Task with the same title already exists"})
			return
		}
		log.Printf("Failed to create task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

func extractBase64FromFile(file *multipart.FileHeader) (string, error) {

	if file.Size > MaxFileSize {
		return "", errors.New("file size exceeds maximum allowed size")
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	data, err := ioutil.ReadAll(src)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, nil
}

// ListTasks handles the retrieval of all tasks.
// @Summary List all tasks
// @Description Retrieve a list of all tasks in the system
// @Accept json
// @Produce json
// @Success 200 {array} models.Task "List of tasks"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /tasks [get]
func (h *TaskHandler) ListTasks(c *gin.Context) {
	tasks, err := h.Service.GetAllTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID handles the retrieval of a task by its ID.
// @Summary Get a task by ID
// @Description Retrieve a task by its unique ID
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} models.Task "Requested task"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Task not found"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	taskID := c.Query("id")

	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing task ID"})
		return
	}

	task, err := h.Service.GetTaskByID(c.Request.Context(), taskID)
	if err != nil {
		log.Printf("Failed to get task by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task by ID"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask handles the updating of a task.
// @Summary Update a task
// @Description Update an existing task in the system
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body models.Task true "Updated task object"
// @Success 200 {object} gin.H "Task updated successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Task not found"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {

	taskID := c.Param("id")
	log.Println("Logging TaskID: ", taskID)

	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Service.UpdateTaskQuery(c.Request.Context(), taskID, &updatedTask)
	if err != nil {
		log.Printf("Failed to update task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTask handles the deletion of a task.
// @Summary Delete a task
// @Description Delete an existing task from the system
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} gin.H "Task deleted successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Task not found"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {

	taskID := c.Param("id")

	err := h.Service.DeleteTaskQuery(c.Request.Context(), taskID)
	if err != nil {
		log.Printf("Failed to delete task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// DeleteAllTasks handles the deletion of all tasks.
// @Summary Delete all tasks
// @Description Delete all existing tasks from the system
// @Accept json
// @Produce json
// @Success 200 {object} gin.H "All tasks deleted successfully"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /tasks [delete]
func (h *TaskHandler) DeleteAllTasks(c *gin.Context) {
	err := h.Service.DeleteAllTasksQuery(c.Request.Context())
	if err != nil {
		log.Printf("Failed to delete all tasks: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all tasks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All tasks deleted successfully"})
}

// MarkTaskAsDone handles marking a task as done.
// @Summary Mark a task as done
// @Description Mark an existing task as done
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} gin.H "Task marked as done successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 404 {object} gin.H "Task not found"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /tasks/done/{id} [put]
func (h *TaskHandler) MarkTaskAsDone(c *gin.Context) {
	taskID := c.Param("id")
	log.Println("Logging TaskID: ", taskID)

	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing task ID"})
		return
	}

	updatedTask, err := h.Service.MarkTaskAsDoneQuery(c.Request.Context(), taskID)
	if err != nil {
		log.Printf("Failed to mark task as done: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark task as done"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task marked as done successfully", "task": updatedTask})
}
