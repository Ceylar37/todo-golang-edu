package repository

import (
	todo "crud"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TodoPostgres struct {
	db *sqlx.DB
}

func NewTodoPostgres(db *sqlx.DB) *TodoPostgres {
	return &TodoPostgres{db: db}
}

func (t *TodoPostgres) GetAllTodos() ([]todo.Todo, error) {
	var todos []todo.Todo
	query := fmt.Sprintf("SELECT id, description, is_done, is_favourite FROM %s", todoTable)
	err := t.db.Select(&todos, query)
	return todos, err
}

func (t *TodoPostgres) CreateTodo(todo todo.Todo) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (description) values ($1) RETURNING id", todoTable)
	row := t.db.QueryRow(query, todo.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (t *TodoPostgres) ChangeDoneStatus(todoData todo.ChangeDoneStatusDto) (int, error) {
	var id int
	query := fmt.Sprintf("UPDATE %s SET is_done = $2 WHERE id = $1 RETURNING id", todoTable)
	row := t.db.QueryRow(query, todoData.Id, todoData.IsDone)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (t *TodoPostgres) ChangeFavouriteStatus(todoData todo.ChangeFavouriteStatusDto) (int, error) {
	var id int
	query := fmt.Sprintf("UPDATE %s SET is_favourite = $2 WHERE id = $1 RETURNING id", todoTable)
	row := t.db.QueryRow(query, todoData.Id, todoData.IsFavourite)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (t *TodoPostgres) DeleteTodo(deleteId int) (int, error) {
	var id int
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", todoTable)
	row := t.db.QueryRow(query, deleteId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
