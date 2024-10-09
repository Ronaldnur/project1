package todo_pg

import (
	"Project1/entity"
	"Project1/pkg/errs"
	"Project1/repository/todo_repository"
	"database/sql"
	"errors"
)

const (
	createTodoData = `
INSERT INTO "data"
(
	title,
	completed
	
)
VALUES ($1, $2)
`

	readTodosData = `
SELECT
	id,
	title,
	completed
FROM "data"
`

	readtodoDataById = `
SELECT
id,
title,
completed
FROM "data"
WHERE id = $1
`

	updateTodoDataById = `
		UPDATE "data"
		SET
			title = $1,
			completed = $2
		WHERE id = $3
	`

	deleteTodoDataById = `
    DELETE FROM "data"
    WHERE id = $1
  `
)

type todoPG struct {
	db *sql.DB
}

func NewTodoPG(db *sql.DB) todo_repository.Repository {
	return &todoPG{
		db: db,
	}
}

func (t *todoPG) CreateTodo(todoPayload entity.Todos) errs.MessageErr {

	tx, err := t.db.Begin()
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(createTodoData, todoPayload.Title, todoPayload.Completed)
	if err != nil {

		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}

func (t *todoPG) ReadTodos() ([]entity.Todos, errs.MessageErr) {
	rows, err := t.db.Query(readTodosData)
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}
	defer rows.Close()

	var todos []entity.Todos
	for rows.Next() {
		var todo entity.Todos
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Completed); err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *todoPG) ReadTodoById(todoId int) (*entity.Todos, errs.MessageErr) {
	var todo entity.Todos
	rows := t.db.QueryRow(readtodoDataById, todoId)

	err := rows.Scan(&todo.Id, &todo.Title, &todo.Completed)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("Todo data not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &todo, nil

}

func (t *todoPG) UpdateTodoById(todoId int, payload entity.Todos) errs.MessageErr {
	_, err := t.db.Exec(updateTodoDataById, payload.Title, payload.Completed, todoId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}

func (t *todoPG) DeleteTodoById(todoId int) errs.MessageErr {
	result, err := t.db.Exec(deleteTodoDataById, todoId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return errs.NewNotFoundError("Todo Id Not Found")
	}
	return nil
}
