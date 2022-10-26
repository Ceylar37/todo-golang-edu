package service

import (
	todo "crud"
	"crud/pkg/repository"
)

type TodoService struct {
	repo repository.Todo
}

func NewTodoService(repo repository.Todo) *TodoService {
	return &TodoService{repo: repo}
}

func (t *TodoService) GetAllTodo() ([]todo.Todo, error) {
	return t.repo.GetAllTodos()
}

func (t *TodoService) CreateTodo(todo todo.Todo) (int, error) {
	return t.repo.CreateTodo(todo)
}

func (t *TodoService) ChangeDoneStatus(todoData todo.ChangeDoneStatusDto) (int, error) {
	return t.repo.ChangeDoneStatus(todoData)
}

func (t *TodoService) ChangeFavouriteStatus(todoData todo.ChangeFavouriteStatusDto) (int, error) {
	return t.repo.ChangeFavouriteStatus(todoData)
}

func (t *TodoService) DeleteTodo(deleteId int) (int, error) {
	return t.repo.DeleteTodo(deleteId)
}
