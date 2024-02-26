package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo_api_gin_golang/api"
	"todo_api_gin_golang/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockTaskService struct {
	Tasks []models.Task
	Err   error
}

func (s *MockTaskService) CreateTaskQuery(ctx context.Context, task *models.Task) (*models.Task, error) {
	return task, nil
}

func (*MockTaskService) DeleteAllTasksQuery(ctx context.Context) error {
	return nil
}

func (*MockTaskService) DeleteTaskQuery(ctx context.Context, id string) error {
	return nil
}

func (m *MockTaskService) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	for _, task := range m.Tasks {
		if task.ID.String() == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (m *MockTaskService) UpdateTaskQuery(ctx context.Context, id string, updatedTask *models.Task) (*models.Task, error) {
	for i, task := range m.Tasks {
		if task.ID.String() == id {
			m.Tasks[i] = *updatedTask
			return updatedTask, nil
		}
	}
	return nil, errors.New("task not found")
}

func (m *MockTaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return m.Tasks, m.Err
}

func (m *MockTaskService) MarkTaskAsDoneQuery(ctx context.Context, id string) (*models.Task, error) {
	for i, task := range m.Tasks {
		if task.ID.String() == id {
			m.Tasks[i].Status = "COMPLETED" // Update task status to "COMPLETED"
			return &m.Tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}


func TestTaskHandler(t *testing.T) {
	t.Run("TestListTasks", TestListTasks)
	t.Run("TestCreateTask", TestCreateTask)
	t.Run("TestGetTaskByID", TestGetTaskByID)
	t.Run("TestUpdateTask", TestUpdateTask)
	t.Run("TestDeleteTask", TestDeleteTask)
	t.Run("TestDeleteAllTask", TestDeleteAllTasks)

	t.Run("TestMarkDoneTask", TestMarkTaskAsDone)
}

func TestListTasks(t *testing.T) {
	mockTasks := []models.Task{
		{ID: uuid.New(), Title: "Task 1"},
		{ID: uuid.New(), Title: "Task 2"},
	}

	mockService := &MockTaskService{
		Tasks: mockTasks,
		Err:   nil,
	}

	handler := api.NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.ListTasks(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseTasks []models.Task
	err = json.Unmarshal(w.Body.Bytes(), &responseTasks)
	assert.Nil(t, err)

	assert.Equal(t, mockTasks, responseTasks)
}

func TestCreateTask(t *testing.T) {
	mockTask := models.Task{ID: uuid.New(), Title: "Task 1"}

	mockService := &MockTaskService{
		Tasks: []models.Task{mockTask},
		Err:   nil,
	}

	handler := api.NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	body, _ := json.Marshal(mockTask)
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.CreateTask(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var responseTask models.Task
	err = json.Unmarshal(w.Body.Bytes(), &responseTask)
	assert.Nil(t, err)

	assert.Equal(t, mockTask, responseTask)
}

func TestGetTaskByID(t *testing.T) {
	mockTask := models.Task{
		ID: uuid.New(),
	}

	mockService := &MockTaskService{
		Tasks: []models.Task{mockTask},
		Err:   nil,
	}

	handler := api.NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	req, err := http.NewRequest(http.MethodGet, "/tasks?id="+mockTask.ID.String(), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.GetTaskByID(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseTask models.Task
	err = json.Unmarshal(w.Body.Bytes(), &responseTask)
	assert.Nil(t, err)

	assert.Equal(t, mockTask, responseTask)
}

func TestUpdateTask(t *testing.T) {

	mockTaskID := uuid.New()
	updatedTask := models.Task{ID: mockTaskID, Title: "Updated Task"}

	mockService := &MockTaskService{
		Tasks: []models.Task{{ID: mockTaskID, Title: "Original Task"}},
		Err:   nil,
	}

	handler := api.NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	body, _ := json.Marshal(updatedTask)
	req, err := http.NewRequest(http.MethodPut, "/tasks/"+mockTaskID.String(), bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: mockTaskID.String()}}

	handler.UpdateTask(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, err)

	expectedMessage := "Task updated successfully"
	assert.Equal(t, expectedMessage, responseBody["message"])
}

func TestDeleteTask(t *testing.T) {

	mockTaskID := uuid.New()

	mockService := &MockTaskService{
		Tasks: nil,
		Err:   nil,
	}

	handler := api.NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	req, err := http.NewRequest(http.MethodDelete, "/tasks/"+mockTaskID.String(), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: mockTaskID.String()}}

	handler.DeleteTask(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, "Task deleted successfully", response["message"])
}

func TestDeleteAllTasks(t *testing.T) {

	mockService := &MockTaskService{
		Err: nil,
	}

	handler := api.NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	req, err := http.NewRequest(http.MethodDelete, "/tasks", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.DeleteAllTasks(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, err)

	expectedMessage := "All tasks deleted successfully"
	assert.Equal(t, expectedMessage, responseBody["message"])
}

func TestMarkTaskAsDone(t *testing.T) {
	mockTaskID := uuid.New()

	mockService := &MockTaskService{
		Tasks: []models.Task{{ID: mockTaskID, Status: "IN_PROGRESS"}},
		Err:   nil,
	}

	handler := api.NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	req, err := http.NewRequest(http.MethodGet, "/tasks/"+mockTaskID.String(), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: mockTaskID.String()}}

	handler.MarkTaskAsDone(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, "COMPLETED", response["task"].(map[string]interface{})["status"]) // Check status instead of "done"
}
