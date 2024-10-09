package handler

import (
	"Project1/dto"
	"Project1/pkg/errs"
	"Project1/pkg/helpers"
	"Project1/service"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	TodoService service.TodoService
}

func NewTodoHandler(TodoService service.TodoService) todoHandler {
	return todoHandler{
		TodoService: TodoService,
	}
}

// GetTodos godoc
// @Tags todos
// @Summary Create a new Todo
// @Description Create a new Todo item.
// @Accept json
// @Produce json
// @Param request body dto.NewTodosRequest true "Todo request body"
// @Success 200 {object} dto.NewTodosResponse
// @Router /todos [post]
func (th *todoHandler) CreateData(ctx *gin.Context) {
	var newTodoRequest dto.NewTodosRequest
	if err := ctx.ShouldBindJSON(&newTodoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	result, err := th.TodoService.CreateTodo(newTodoRequest)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

// GetTodos godoc
// @Tags todos
// @Summary Get Todos Data
// @Description Get Todos Data
// @ID get-Todo-data
// @Produce json
// @Success 200 {object} dto.ReadDataTodoResponse
// @Router /todos [get]
func (th *todoHandler) ReadData(ctx *gin.Context) {
	result, err := th.TodoService.ReadTodos()
	if err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

// GetTodoById godoc
// @Tags todos
// @Summary Get Todo Data by ID
// @Description Get a Todo by ID
// @ID get-Todo-by-ID
// @Produce json
// @Param todoId path int true "Todo ID"
// @Success 200 {object} dto.ReadDataByIdTodoResponse
// @Router /todos/{todoId} [get]
func (th *todoHandler) ReadDataById(ctx *gin.Context) {
	todoID, err := helpers.GetParamId(ctx, "todoId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := th.TodoService.ReadTodoById(todoID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}
	ctx.JSON(result.StatusCode, result)
}

// UpdateTodoById godoc
// @Tags todos
// @Summary Update Todo Data
// @Description Update a Todo by ID
// @ID update-Todo-by-ID
// @Accept json
// @Produce json
// @Param todoId path int true "Todo ID"
// @Param request body dto.NewTodosRequest true "Todo request body"
// @Success 200 {object} dto.NewTodosResponse
// @Router /todos/{todoId} [put]
func (th *todoHandler) UpdateDataById(ctx *gin.Context) {
	var todoRequest dto.NewTodosRequest

	if err := ctx.ShouldBindJSON(&todoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	todoId, err := helpers.GetParamId(ctx, "todoId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	result, err := th.TodoService.UpdateTodoById(todoId, todoRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

// DeleteTodoById godoc
// @Tags todos
// @Summary Delete Todo Data
// @Description Delete a Todo by ID
// @ID delete-Todo-by-ID
// @Produce json
// @Param todoId path int true "Todo ID"
// @Success 204 "No Content"
// @Router /todos/{todoId} [delete]
func (th *todoHandler) DeleteTodoData(ctx *gin.Context) {
	todoID, err := helpers.GetParamId(ctx, "todoId")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := th.TodoService.DeleteTodoById(todoID)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}
