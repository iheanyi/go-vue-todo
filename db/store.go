package db

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

type Store struct {
	db *bolt.DB
}

func (s *Store) CreateTodo(todo *Todo) (*Todo, error) {
	todo.CreatedAt = time.Now().Format(time.RFC3339)
	todo.UpdatedAt = time.Now().Format(time.RFC3339)
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))

		id, _ := b.NextSequence()
		todo.Id = int(id)

		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}

		return b.Put(itob(todo.Id), buf)
	})
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Store) EditTodo(id int, t *Todo) (*Todo, error) {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))

		t.UpdatedAt = time.Now().Format(time.RFC3339)
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}
		return b.Put(itob(id), buf)
	})
	if err != nil {
		return nil, err
	}

	var todo *Todo
	err = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))
		v := b.Get(itob(id))
		if err := json.Unmarshal(v, todo); err != nil {
			return err
		}

		return nil
	})

	return todo, err
}

func (s *Store) DeleteTodo(id int) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))

		return b.Delete(itob(id))
	})

	return err
}

func (s *Store) ListTodos() ([]*Todo, error) {
	var todos []*Todo = make([]*Todo, 0)
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))
		err := b.ForEach(func(k, v []byte) error {
			var t Todo
			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}
			todos = append(todos, &t)
			return err
		})
		return err
	})

	return todos, err
}

func NewStore(db *bolt.DB) *Store {
	return &Store{
		db: db,
	}
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
