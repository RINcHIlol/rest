package service

import (
	restApi "github.com/RINcHIlol/rest.git"
	"github.com/RINcHIlol/rest.git/pkg/repository"
)

//service - работа с бизнес логикой

type Authorization interface {
	CreateUser(user restApi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list restApi.TodoList) (int, error)
	GetAll(userId int) ([]restApi.TodoList, error)
	GetById(userId, id int) (restApi.TodoList, error)
	Delete(userId, id int) error
	Update(userId, id int, input restApi.UpdateTodoList) error
}

type TodoItem interface {
	Create(userId, listId int, item restApi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]restApi.TodoItem, error)
	GetById(userId, itemId int) (restApi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, id int, input restApi.UpdateTodoItem) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
