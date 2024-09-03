package repository

//repository - работа с бд

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
