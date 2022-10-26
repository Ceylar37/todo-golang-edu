package repository

import (
	todo "crud"
	"github.com/jmoiron/sqlx"
)

type Todo interface {
	GetAllTodos() ([]todo.Todo, error)
	CreateTodo(todo todo.Todo) (int, error)
	ChangeDoneStatus(todoData todo.ChangeDoneStatusDto) (int, error)
	ChangeFavouriteStatus(todoData todo.ChangeFavouriteStatusDto) (int, error)
	DeleteTodo(deleteId int) (int, error)
}

type Repository struct {
	Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Todo: NewTodoPostgres(db),
	}
}
