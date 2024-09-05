package repository

import (
	restApi "github.com/RINcHIlol/rest.git"
	"github.com/jmoiron/sqlx"
)

//repository - работа с бд

type Authorization interface {
	CreateUser(user restApi.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
