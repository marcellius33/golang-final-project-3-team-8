package repositories

import (
	"kanbanboard/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) (*models.Task, error)
	GetTasks(userId uint, tasks *[]models.Task) (*[]models.Task, error)
	UpdateTask(taskId uint, task *models.Task) (*models.Task, error)
	DeleteTask(id uint) error
	FindTaskByUserId(askId uint, userId uint) (*models.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (c *taskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	return task, c.db.Create(task).Error
}

func (c *taskRepository) GetTasks(userId uint,tasks *[]models.Task) (*[]models.Task, error) {
	err := c.db.Model(&models.Task{}).Where("user_id=?", userId).Find(&tasks).Error
	return tasks, err
}

func (c *taskRepository) UpdateTask(taskId uint, updateTask *models.Task) (*models.Task, error) {
	task := updateTask
	err := c.db.Model(&task).Where("id=?", taskId).Updates(updateTask).Error
	return task, err
}

func (c *taskRepository) DeleteTask(taskId uint) error {
	err := c.db.Where("id=?", taskId).Delete(&models.Task{}).Error
	return err
}

func (c *taskRepository) FindTaskByUserId(taskId uint, userId uint) (*models.Task, error) {
	task := models.Task{}
	err := c.db.Where("id=? AND user_id=?", taskId, userId).First(&task).Error
	return &task, err
}
