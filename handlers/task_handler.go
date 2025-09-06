package handlers

import (
	"net/http"
	"strconv"

	"github.com/aluyapeter/coco/models"
	"github.com/aluyapeter/coco/repository"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	repo  repository.TaskRepository
}

func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{
		repo: repo,
	}
}

// creating a new task POST /api/v1/tasks
func (h *TaskHandler) CreatetTask(c *gin.Context) {
	var req models.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	task := models.Task{
		Title: req.Title,
		Description: req.Description,
	}

	//save task to database
	if err := h.repo.Create(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create task",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"data": task.ToResponse(),
	})
}

//Get tasks GET /api/v1/tasks
func (h *TaskHandler) GetTasks(c *gin.Context) {
	status := c.Query("status")

	var tasks []models.Task
	var err error

	//get tasks based on filter
	switch status {
	case "completed":
		tasks, err = h.repo.GetCompleted()
	case "pending":
		tasks, err = h.repo.GetPending()
	default:
		tasks, err = h.repo.GetAll()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retreive tasks",
		})
		return
	}

	var responses []models.TaskResponse
	for _, task := range tasks {
		responses = append(responses, task.ToResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responses,
		"count": len(responses),
	})
}

// get tasks by id GET /api/v1/tasks/:id

func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task id",
		})
		return
	}

	//get tasks from database
	task, err := h.repo.GetByID(uint(id))
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Task not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve task",
		})
		return
	}

	// return task
	c.JSON(http.StatusOK, gin.H{
		"data": task.ToResponse(),
	})
}

//update tasks PUT /api/v1/tasks/:id
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task by ID",
		})
		return
	}

	var req models.UpdateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	updates := models.Task{}
	if req.Title != nil {
		updates.Title = *req.Title
	}
	if req.Description != nil {
		updates.Description = *req.Description
	}
	if req.Completed != nil {
		updates.Completed = *req.Completed
	}

	// update in database
	task, err := h.repo.Update(uint(id), &updates)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Task not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update task",
		})
		return
	}

	// return updated task
	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"data": task.ToResponse(),
	})
}

//Delete task by id DELETE /api/v1/tasks/:id
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	//delete from database
	if err := h.repo.Delete(uint(id)); err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Task not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete task",
		})
		return
	}

	// return success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}

// gettaskstats{ Returns statistics about tasks (total, completed, pending)} GET /api/v1/tasks/stats 
func (h *TaskHandler) GetTaskStats(c *gin.Context) {
	//getting all tasks

	allTasks, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve task statistics",
		})
		return
	}

	//get completed tasks
	completedTasks, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve task statistics",
		})
		return
	}

	//get pending tasks
	pendingTasks, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve task statistics",
		})
		return
	}

	//return stats
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"total":     len(allTasks),
			"completed": len(completedTasks),
			"pending":   len(pendingTasks),
		},
	})
}