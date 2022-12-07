package controllers

import (
	"kanbanboard/helpers"
	"kanbanboard/params"
	"kanbanboard/services"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service services.TaskService
}

func NewTaskController(service services.TaskService) *TaskController {
	return &TaskController{
		service: service,
	}
}

func (p *TaskController) CreateTask(c *gin.Context) {
	TaskRequest := params.CreateTaskRequest{}
	if err := c.ShouldBindJSON(&TaskRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	userData, _ := c.MustGet("userData").(jwt.MapClaims)

	createTask, err := p.service.CreateTask(uint(userData["id"].(float64)), TaskRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessCreateResponse(createTask, "Create Task Success"))
}

func (p *TaskController) GetTasks(c *gin.Context) {
	userData, _ := c.MustGet("userData").(jwt.MapClaims)
	Tasks, err := p.service.GetTasks(uint(userData["id"].(float64)))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(Tasks, "Get Tasks Success"))
}

func (p *TaskController) UpdateTask(c *gin.Context) {
	taskRequest := params.UpdateTaskRequest{}
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	userData, _ := c.MustGet("userData").(jwt.MapClaims)

	updateTask, err := p.service.UpdateTask(uint(taskId), uint(userData["id"].(float64)), taskRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(updateTask, "Update Task Success"))
}

func (p *TaskController) UpdateTaskStatus(c *gin.Context) {
	taskStatusRequest := params.UpdateTaskStatusRequest{}
	if err := c.ShouldBindJSON(&taskStatusRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	userData, _ := c.MustGet("userData").(jwt.MapClaims)

	updateTask, err := p.service.UpdateTaskStatus(uint(taskId), uint(userData["id"].(float64)), taskStatusRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(updateTask, "Update Task Success"))
}

func (p *TaskController) UpdateTaskCategory(c *gin.Context) {
	taskCategoryRequest := params.UpdateTaskCategoryRequest{}
	if err := c.ShouldBindJSON(&taskCategoryRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	userData, _ := c.MustGet("userData").(jwt.MapClaims)

	updateTask, err := p.service.UpdateTaskCategory(uint(taskId), uint(userData["id"].(float64)), taskCategoryRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(updateTask, "Update Task Success"))
}

func (p *TaskController) DeleteTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	userData, _ := c.MustGet("userData").(jwt.MapClaims)
	err = p.service.DeleteTask(uint(taskId), uint(userData["id"].(float64)))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.DeleteSuccess("Your task has been successfully deleted"))
}
