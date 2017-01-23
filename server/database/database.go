package database

import (
	pb "github.com/iheanyi/go-electron-grpc/demo"
)

type Database interface {
	CreateTodo(todo *pb.Todo) (*pb.Todo, error)
}
