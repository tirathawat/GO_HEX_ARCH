package controller

import (
	"todo/errs"
	"todo/logs"
	"todo/models"
	"todo/repository"
	"todo/utils"
)

type todoController struct {
	repository repository.ITodoRepository
}

type ITodoController interface {
	GetAll() ([]models.TodoResponse, error)
}

func NewTodoController(todoRepository repository.ITodoRepository) todoController {
	return todoController{repository: todoRepository}
}

func (c todoController) GetAll() ([]models.TodoResponse, error) {
	todoResponses := make([]models.TodoResponse, 0)

	todo, err := c.repository.GetAll()
	if err != nil {
		logs.New().Error(err)
		return nil, errs.NewNotFoundError("ไม่พบข้อมูล", "Items Not Found")
	}

	err = utils.NewType().StructToStruct(todo, &todoResponses)
	if err != nil {
		logs.New().Error(err)
		return nil, err
	}
	return todoResponses, err
}
