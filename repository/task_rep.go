package repository

import (
	"errors"

	"github.com/aluyapeter/coco/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]models.Task, error)
	GetByID(id uint) (*models.Task, error)
	Update(id uint, updates *models.Task) (*models.Task, error)
	Delete(id uint) error
	GetCompleted() ([]models.Task, error)
	GetPending() ([]models.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository {
		db: db,
	}
}

func (r *taskRepository) Create(task *models.Task) error {
	if err := r.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// Retrieve all tasks form the database
func (r *taskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	err := r.db.Order("created_at DESC").Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetByID retrieves a single task by its Id from the database
func (r *taskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task

	//find by ID
	err := r.db.First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return &task, nil
}

// Update an existing task
func (r *taskRepository) Update(id uint, updates *models.Task) (*models.Task, error) {
	var task models.Task

	//check if task exists
	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	// update task with new values
	if err := r.db.Model(&task).Updates(updates).Error; err != nil {
		return nil, err
	}

	// get updated task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

//deleting a task from database

func (r *taskRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Task{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

//get completed task
func (r *taskRepository) GetCompleted() ([]models.Task, error) {
	var tasks []models.Task

	err := r.db.Where("completed = ?", true).Order("created_at DESC").Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// get pending/not completed tasks
func (r *taskRepository) GetPending() ([]models.Task, error) {
	var tasks []models.Task

	err := r.db.Where("completed = ?", false).Order("created_at DESC").Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}