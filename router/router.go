package router

import (
	"todo/controller"
	"todo/database"
	"todo/handler"
	"todo/repository"

	"github.com/labstack/echo/v4"
)

type router struct {
	echo *echo.Echo
}

type IRouter interface {
}

var instantiated *router = nil

func New(e *echo.Echo) *router {
	if instantiated == nil {
		instantiated = &router{echo: e}
		instantiated.init()
	}
	return instantiated
}

func (r router) init() {
	r.setupTodo()
}

func (r router) setupTodo() {
	db := database.New().GetDB()
	repository := repository.NewTodoRepository(db)
	controller := controller.NewTodoController(repository)
	handler := handler.NewTodoHandler(controller)
	group := r.echo.Group("todo")
	{
		group.GET("/list", handler.GetAll)
	}
}
