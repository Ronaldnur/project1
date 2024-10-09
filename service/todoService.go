package service

import (
	"Project1/dto"
	"Project1/entity"
	"Project1/pkg/errs"
	"Project1/repository/todo_repository"
	"net/http"
)

type todoService struct {
	TodoRepo todo_repository.Repository
}

type TodoService interface {
	CreateTodo(todoPayload dto.NewTodosRequest) (*dto.NewTodosResponse, errs.MessageErr)
	ReadTodos() (*dto.ReadDataTodoResponse, errs.MessageErr)
	ReadTodoById(todoId int) (*dto.ReadDataByIdTodoResponse, errs.MessageErr)
	UpdateTodoById(todoId int, payload dto.NewTodosRequest) (*dto.NewTodosResponse, errs.MessageErr)
	DeleteTodoById(todoId int) (*dto.NewTodosResponse, errs.MessageErr)
}

func NewTodoService(TodoRepo todo_repository.Repository) TodoService {
	return &todoService{
		TodoRepo: TodoRepo,
	}
}
func (t *todoService) CreateTodo(todoPayload dto.NewTodosRequest) (*dto.NewTodosResponse, errs.MessageErr) {
	TodoRequest := entity.Todos{
		Title:     todoPayload.Title,
		Completed: todoPayload.Completed,
	}
	err := t.TodoRepo.CreateTodo(TodoRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewTodosResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully create Todo Data",
		Result:     "Success",
	}
	return &response, nil
}

func (t *todoService) ReadTodos() (*dto.ReadDataTodoResponse, errs.MessageErr) {
	Todos, err := t.TodoRepo.ReadTodos()
	if err != nil {
		return nil, err
	}

	todoResult := []dto.TodoWithId{}

	for _, eachTodo := range Todos {
		Todo := dto.TodoWithId{
			Id:        eachTodo.Id,
			Title:     eachTodo.Title,
			Completed: eachTodo.Completed,
		}
		todoResult = append(todoResult, Todo)
	}
	response := dto.ReadDataTodoResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfuly Read Todo Data",
		Todos:      todoResult,
	}

	return &response, nil
}

func (t *todoService) ReadTodoById(todoId int) (*dto.ReadDataByIdTodoResponse, errs.MessageErr) {
	Todo, err := t.TodoRepo.ReadTodoById(todoId)
	if err != nil {
		return nil, err
	}

	response := dto.ReadDataByIdTodoResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully Read Todo Data by ID",
		Todo: dto.TodoWithId{
			Id:        Todo.Id,
			Title:     Todo.Title,
			Completed: Todo.Completed,
		},
	}
	return &response, nil
}

func (t *todoService) UpdateTodoById(TodoId int, payload dto.NewTodosRequest) (*dto.NewTodosResponse, errs.MessageErr) {

	_, err := t.TodoRepo.ReadTodoById(TodoId)
	if err != nil {
		return nil, errs.NewNotFoundError("Todo Id Not Found")
	}

	TodoRequest := entity.Todos{
		Title:     payload.Title,
		Completed: payload.Completed,
	}
	err = t.TodoRepo.UpdateTodoById(TodoId, TodoRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewTodosResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully updated Todo Data",
		Result:     "Success",
	}

	return &response, nil
}

func (t *todoService) DeleteTodoById(TodoId int) (*dto.NewTodosResponse, errs.MessageErr) {
	err := t.TodoRepo.DeleteTodoById(TodoId)
	if err != nil {
		return nil, err
	}

	result := dto.NewTodosResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully deleted Todo Data",
		Result:     "Success",
	}

	return &result, nil
}
