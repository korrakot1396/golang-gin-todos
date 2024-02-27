package models

import (
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id" swagger:"description=Unique identifier for the task"`
	Title       string    `json:"title" swagger:"description=Title of the task"`
	Date        string    `json:"date" swagger:"description=Date of the task"`
	Status      string    `json:"status" swagger:"description=Status of the task"`
	Description string    `json:"description" swagger:"description=Description of the task"`
	Image       string    `json:"image,omitempty" swagger:"description=Base64 encoded image of the task, if available"`
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
