package database

import (
	"github.com/boltdb/bolt"
	pb "github.com/iheanyi/go-electron-grpc/demo"
	"os"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	db, teardown, err := setupTestDatabase()
	if err != nil {
		t.Errorf("Error creating test database: %v", err)
	}
	defer teardown()

	store := NewStore(db)
	mock := &pb.Todo{
		Description: "Test this",
		Done:        false,
	}

	output, err := store.CreateTodo(mock)
	if err != nil {
		t.Errorf("Error thrown by CreateTodo: %v", err)
	}

	if output.Description != mock.Description {
		t.Errorf("Got %v, want %v", output.Description, mock.Description)
	}

	if output.Done != mock.Done {
		t.Errorf("Got %v, want %v", output.Done, mock.Done)
	}

	if output.Id < 1 {
		t.Errorf("Expected Todo.Id to not be zero.")
	}
}

func TestListTodos(t *testing.T) {
	db, teardown, err := setupTestDatabase()
	if err != nil {
		t.Errorf("Error setting up test database: %v", err)
	}
	defer teardown()

	store := NewStore(db)

	store.CreateTodo(&pb.Todo{
		Description: "Test this.",
		Done:        false,
	})

	store.CreateTodo(&pb.Todo{
		Description: "Another test.",
		Done:        false,
	})

	store.CreateTodo(&pb.Todo{
		Description: "Last test.",
		Done:        false,
	})

	output, err := store.ListTodos()
	if err != nil {
		t.Errorf("store.ListTodos returned err: %v", err)
	}

	if len(output) != 3 {
		t.Errorf("Expected output to be of length 3, got length: %v", len(output))
	}
}

func setupTestDatabase() (*bolt.DB, func(), error) {
	// Wipe the database before every single test.
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		return nil, nil, err
	}

	return db, func() {
		db.Close()
		os.Remove("test.db")
	}, nil
}
