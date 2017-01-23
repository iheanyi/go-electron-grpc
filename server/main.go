package main

import (
	"github.com/boltdb/bolt"
	pb "github.com/iheanyi/go-electron-grpc/demo"
	"github.com/iheanyi/go-electron-grpc/server/database"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	store database.Database
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Print("Request received!")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	db, err := bolt.Open("demo.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	store := database.NewStore(db)
	pb.RegisterGreeterServer(s, &server{
		store: store,
	})

	reflection.Register(s)
	log.Print("Starting up the Go Server.")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
