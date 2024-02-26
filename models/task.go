package models

import (
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Date        string    `json:"date"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Image       string    `json:"imgBase64,omitempty"`
}

func NewTask(title, date, status, description, image string) *Task {
	return &Task{
		ID:          uuid.New(),
		Title:       title,
		Date:        date,
		Status:      status,
		Description: description,
		Image:       image,
	}
}
