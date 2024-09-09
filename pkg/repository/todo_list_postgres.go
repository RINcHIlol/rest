package repository

import (
	"fmt"
	restApi "github.com/RINcHIlol/rest.git"
	"github.com/jmoiron/sqlx"
)

type ToDoListPostgres struct {
	db *sqlx.DB
}

func NewToDoListPostgres(db *sqlx.DB) *ToDoListPostgres {
	return &ToDoListPostgres{db: db}
}

func (s *ToDoListPostgres) Create(userId int, list restApi.TodoList) (int, error) {
	var id int

	tx, err := s.db.Begin()
	if err != nil {
		return 0, nil
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	CreateUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(CreateUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (s *ToDoListPostgres) GetAll(userId int) ([]restApi.TodoList, error) {
	var lists []restApi.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := s.db.Select(&lists, query, userId)
	return lists, err
}

func (s *ToDoListPostgres) GetById(userId, id int) (restApi.TodoList, error) {
	var list restApi.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl"+
		" INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 and ul.list_id = $2",
		todoListsTable, usersListsTable)
	err := s.db.Get(&list, query, userId, id)
	return list, err
}
