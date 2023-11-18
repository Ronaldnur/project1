package todo_repository

import (
	"Project1/entity"
	"Project1/pkg/errs"
)

type Repository interface {
	CreateTodo(todoPayload entity.Todos) errs.MessageErr
	ReadTodos() ([]entity.Todos, errs.MessageErr)
	ReadTodoById(todoId int) (*entity.Todos, errs.MessageErr)
	UpdateTodoById(todoId int, payload entity.Todos) errs.MessageErr
	DeleteTodoById(todoId int) errs.MessageErr
}
