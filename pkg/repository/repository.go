package repository

import (
	restApi "github.com/RINcHIlol/rest.git"
	"github.com/jmoiron/sqlx"
)

//repository - работа с бд

type Authorization interface {
	CreateUser(user restApi.User) (int, error)
	GetUser(username, password string) (restApi.User, error)
}

type TodoList interface {
	Create(userId int, list restApi.TodoList) (int, error)
	GetAll(userId int) ([]restApi.TodoList, error)
	GetById(userId, id int) (restApi.TodoList, error)
	Delete(userId, id int) error
	Update(userId, id int, input restApi.UpdateTodoList) error
}

type TodoItem interface {
	Create(listId int, item restApi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]restApi.TodoItem, error)
	GetById(userId, itemId int) (restApi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, id int, input restApi.UpdateTodoItem) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewToDoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
