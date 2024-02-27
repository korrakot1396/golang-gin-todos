package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"todo_api_gin_golang/models"

	"github.com/google/uuid"
)

type TaskServiceInterface interface {
	CreateTaskQuery(ctx context.Context, task *models.Task) (*models.Task, error)
	GetAllTasks(ctx context.Context) ([]models.Task, error)
	GetTaskByID(ctx context.Context, id string) (*models.Task, error)
	UpdateTaskQuery(ctx context.Context, id string, updatedTask *models.Task) (*models.Task, error)
	DeleteTaskQuery(ctx context.Context, id string) error
	DeleteAllTasksQuery(ctx context.Context) error
	MarkTaskAsDoneQuery(ctx context.Context, id string) (*models.Task, error)
}

type TaskService struct {
	db *sql.DB
}

func NewTaskService(dbPath string) (*TaskService, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
        id TEXT PRIMARY KEY,
        title TEXT,
        date TEXT,
        status TEXT,
		description TEXT,
        image TEXT
    )`)

	if err != nil {
		return nil, fmt.Errorf("failed to create todos table: %v", err)
	}

	return &TaskService{db: db}, nil
}

func (s *TaskService) Close() {
	s.db.Close()
}

func (s *TaskService) CreateTaskQuery(ctx context.Context, task *models.Task) (*models.Task, error) {
	var count int
	err := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM todos WHERE title = ?", task.Title).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("failed to check for duplicate task: %v", err)
	}
	if count > 0 {
		return nil, errors.New("task with the same title already exists")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}
	task.ID = id
	task.Date = time.Now().Format(time.RFC3339)
	task.Status = "IN_PROGRESS"

	_, err = s.db.ExecContext(ctx, "INSERT INTO todos (id, title, description, image, date, status) VALUES (?, ?, ?, ?, ?, ?)", task.ID.String(), task.Title, task.Description, task.Image, task.Date, task.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to insert task: %v", err)
	}

	return task, nil
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {

	query := "SELECT id, title, date, status, image FROM todos ORDER BY date DESC"

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %v", err)
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {

		var task models.Task

		if err := rows.Scan(&task.ID, &task.Title, &task.Date, &task.Status, &task.Image); err != nil {
			return nil, fmt.Errorf("failed to scan task row: %v", err)
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %v", err)
	}

	return tasks, nil
}

func (s *TaskService) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	var task models.Task
	query := "SELECT id, title, date, status, image FROM todos WHERE id = ?"
	err := s.db.QueryRowContext(ctx, query, id).Scan(&task.ID, &task.Title, &task.Date, &task.Status, &task.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task with ID %s not found", id)
		}
		return nil, fmt.Errorf("failed to get task: %v", err)
	}
	return &task, nil
}

func (s *TaskService) UpdateTaskQuery(ctx context.Context, id string, updatedTask *models.Task) (*models.Task, error) {
	existingTask, err := s.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if updatedTask.Title != "" {
		existingTask.Title = updatedTask.Title
	}
	if updatedTask.Date != "" {
		existingTask.Date = updatedTask.Date
	}
	if updatedTask.Image != "" {
		existingTask.Image = updatedTask.Image
	}
	if updatedTask.Status != "" {
		existingTask.Status = updatedTask.Status
	}

	if _, err := s.db.ExecContext(ctx, `
        UPDATE todos
        SET title = ?, date = ?, image = ?, status = ?
        WHERE id = ?
    `, existingTask.Title, existingTask.Date, existingTask.Image, existingTask.Status, id); err != nil {
		return nil, err
	}

	return existingTask, nil
}

func (s *TaskService) DeleteTaskQuery(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) DeleteAllTasksQuery(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM todos")
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) MarkTaskAsDoneQuery(ctx context.Context, id string) (*models.Task, error) {

	existingTask, err := s.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	existingTask.Status = "COMPLETED"

	_, err = s.db.ExecContext(ctx, "UPDATE todos SET status = ? WHERE id = ?", "COMPLETED", id)
	if err != nil {
		return nil, err
	}

	return existingTask, nil
}
