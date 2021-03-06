package db

type Database interface {
	CreateTodo(*Todo) (*Todo, error)
	ListTodos() ([]*Todo, error)
	DeleteTodo(id int) error
	UpdateTodo(id int, t *Todo) (*Todo, error)
}
