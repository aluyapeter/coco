package models

import (
	"time"

	"gorm.io/gorm"
)

//This struct represents the structure of a task in our database
type Task struct {
ID uint `gorm:"primaryKey" json:"id"`
Title string `gorm:"not null;size:255" json:"title" binding:"required"`
Description string `gorm:"type:text" json:"description"`
Completed bool `gorm:"default:false" json:"completed"`
CreatedAt time.Time `json:"created_at"`
UpdatedAt time.Time `json:"updated_at"`
DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

//This struct represents the data that is needed for a user to create a new task
type CreateTaskRequest struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
	Description string `json:"description,omitempty"`
}

//This struct represents the data that is needed for a user to update an existing task
type UpdateTaskRequest struct {
	Title *string `json:"title,omitempty" binding:"omitempty,min=1,max=255"`
	Description *string `json:"description,omitempty"`
	Completed *bool `json:"completed,omitempty"`
}

//This struct represents how data is sent back to the users
type TaskResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Task) ToResponse() TaskResponse{
	return TaskResponse{
		ID: t.ID,
		Title: t.Title,
		Description: t.Description,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func (Task) TableName() string {
	return "tasks"
}