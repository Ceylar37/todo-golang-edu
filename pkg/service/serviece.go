package service

import (
	todo "crud"
	"crud/pkg/repository"
)

type Todo interface {
	GetAllTodo() ([]todo.Todo, error)
	CreateTodo(todo todo.Todo) (int, error)
	ChangeDoneStatus(todoData todo.ChangeDoneStatusDto) (int, error)
	ChangeFavouriteStatus(todoData todo.ChangeFavouriteStatusDto) (int, error)
	DeleteTodo(deleteId int) (int, error)
}

type Service struct {
	Todo
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Todo: NewTodoService(repository.Todo),
	}
}
