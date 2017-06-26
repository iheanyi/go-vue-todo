package db

import (
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) (store *Store, cleanup func()) {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}

	tx, err := db.Begin(true)
	if err != nil {
		t.Fatalf("error beginning tx: %v", err)
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte("todos"))
	if err != nil {
		t.Fatalf("error creating bucket: %v", err)
	}

	if err = tx.Commit(); err != nil {
		t.Fatalf("error commiting tx: %v", err)
	}

	store = NewStore(db)

	return store, func() {
		os.Remove("test.db")
		db.Close()
	}
}

func TestEditTodo(t *testing.T) {
	tests := []struct {
		desc        string
		update      *Todo
		expected    *Todo
		errExpected bool
		setupTest   func(t *testing.T, store *Store) error
	}{}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			store, cleanup := setupTestDB(t)
			defer cleanup()

			if tt.setupTest != nil {
				tt.setupTest(t, store)
			}

			todo, err := store.EditTodo(tt.update.Id, tt.update)
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, todo.Id)
				assert.Equal(t, tt.expected, todo)
			}
		})
	}
}
