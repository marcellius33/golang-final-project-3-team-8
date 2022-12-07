package services

import (
	"errors"
	"kanbanboard/models"
	"kanbanboard/params"
	"kanbanboard/repositories"
)

type TaskService interface {
	CreateTask(userId uint, createTaskRequest params.CreateTaskRequest) (*params.CreateTaskResponse, error)
	GetTasks(userId uint) (*[]params.GetTaskResponse, error)
	UpdateTask(taskId uint, userId uint, taskUpdateRequest params.UpdateTaskRequest) (*params.UpdateTaskResponse, error)
	UpdateTaskStatus(taskId uint, userId uint, taskUpdateRequest params.UpdateTaskStatusRequest) (*params.UpdateTaskResponse, error)
	UpdateTaskCategory(taskId uint, userId uint, taskUpdateRequest params.UpdateTaskCategoryRequest) (*params.UpdateTaskResponse, error)
	DeleteTask(taskId uint, userId uint) error
}

type taskService struct {
	taskR repositories.TaskRepository
	categoryR repositories.CategoryRepository
	userR    repositories.UserRepository
}

func NewTaskService(taskR repositories.TaskRepository, categoryR repositories.CategoryRepository, userR repositories.UserRepository) TaskService {
	return &taskService{
		taskR: taskR,
		categoryR: categoryR,
		userR:    userR,
	}
}

func (c *taskService) CreateTask(userId uint, createTaskRequest params.CreateTaskRequest) (*params.CreateTaskResponse, error) {
	_, err := c.categoryR.FindCategoryById(createTaskRequest.CategoryID)

	if err != nil {
		return &params.CreateTaskResponse{}, err
	}
	
	newTask := models.Task{
		Title: createTaskRequest.Title,
		Description: createTaskRequest.Description,
		UserID: userId,
		CategoryID: createTaskRequest.CategoryID,
	}

	_, err = c.taskR.CreateTask(&newTask)

	if err != nil {
		return &params.CreateTaskResponse{}, err
	}
	resp := params.ParseToCreateTaskResponse(&newTask)

	return &resp, nil
}

func (c *taskService) GetTasks(userId uint) (*[]params.GetTaskResponse, error) {
	var tasks []models.Task
	_, err := c.taskR.GetTasks(userId, &tasks)

	if err != nil {
		return &[]params.GetTaskResponse{}, err
	}
	user, _ := c.userR.FindUserByID(userId)
	resp := params.ParseToGetTasksResponse(tasks, *user)

	return &resp, nil
}

func (c *taskService) UpdateTask(taskId uint, userId uint, taskUpdateRequest params.UpdateTaskRequest) (*params.UpdateTaskResponse, error) {
	taskModel, err := c.taskR.FindTaskByUserId(taskId, userId)
	if err != nil {
		return &params.UpdateTaskResponse{}, errors.New("Unauthorized")
	}
	taskModel.Title = taskUpdateRequest.Title
	taskModel.Description = taskUpdateRequest.Description

	_, err = c.taskR.UpdateTask(taskId, taskModel)

	if err != nil {
		return &params.UpdateTaskResponse{}, err
	}
	resp := params.ParseToUpdateTaskResponse(taskModel)

	return &resp, nil
}

func (c *taskService) UpdateTaskCategory(taskId uint, userId uint, taskUpdateRequest params.UpdateTaskCategoryRequest) (*params.UpdateTaskResponse, error) {
	_, err := c.categoryR.FindCategoryById(taskUpdateRequest.CategoryID)

	if err != nil {
		return &params.UpdateTaskResponse{}, err
	}
	
	taskModel, err := c.taskR.FindTaskByUserId(taskId, userId)
	if err != nil {
		return &params.UpdateTaskResponse{}, errors.New("Unauthorized")
	}
	taskModel.CategoryID = taskUpdateRequest.CategoryID

	_, err = c.taskR.UpdateTask(taskId, taskModel)

	if err != nil {
		return &params.UpdateTaskResponse{}, err
	}
	resp := params.ParseToUpdateTaskResponse(taskModel)

	return &resp, nil
}

func (c *taskService) UpdateTaskStatus(taskId uint, userId uint, taskUpdateRequest params.UpdateTaskStatusRequest) (*params.UpdateTaskResponse, error) {
	taskModel, err := c.taskR.FindTaskByUserId(taskId, userId)
	if err != nil {
		return &params.UpdateTaskResponse{}, errors.New("Unauthorized")
	}
	taskModel.Status= taskUpdateRequest.Status

	_, err = c.taskR.UpdateTask(taskId, taskModel)

	if err != nil {
		return &params.UpdateTaskResponse{}, err
	}
	resp := params.ParseToUpdateTaskResponse(taskModel)

	return &resp, nil
}

func (c *taskService) DeleteTask(taskId uint, userId uint) error {
	_, err := c.taskR.FindTaskByUserId(taskId, userId)
	if err != nil {
		return errors.New("Unauthorized")
	}
	return c.taskR.DeleteTask(taskId)
}

