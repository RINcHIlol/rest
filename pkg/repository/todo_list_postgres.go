package repository

import (
	"fmt"
	restApi "github.com/RINcHIlol/rest.git"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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

func (s *ToDoListPostgres) Delete(userid, id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id and ul.user_id=$1 and ul.list_id=$2",
		todoListsTable, usersListsTable)
	_, err := s.db.Exec(query, userid, id)
	return err
}

func (s *ToDoListPostgres) Update(userId, id int, input restApi.UpdateTodoList) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, id, userId)

	logrus.Debugf("updateQuery s", query)
	logrus.Debugf("args: %s", args)

	_, err := s.db.Exec(query, args...)
	return err
}
