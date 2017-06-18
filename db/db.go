package db

type Database interface {
	CreateTodo(*Todo) (*Todo, error)
	ListTodos() ([]*Todo, error)
}
