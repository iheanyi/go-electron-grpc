package database

import (
	"encoding/binary"
	"encoding/json"
	"github.com/boltdb/bolt"
	pb "github.com/iheanyi/go-electron-grpc/demo"
	"log"
)

type Store struct {
	db *bolt.DB
}

func (s *Store) CreateTodo(todo *pb.Todo) (*pb.Todo, error) {
	log.Print("Entered the CreateTodo method with Todo: %v", todo)

	err := s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("todos"))
		if err != nil {
			return err
		}

		id, _ := b.NextSequence()
		todo.Id = id

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

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func NewStore(db *bolt.DB) *Store {
	return &Store{
		db: db,
	}
}
