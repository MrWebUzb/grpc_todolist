package repo

import (
	pb "github.com/MrWebUzb/microservice/todo_service/service/todo/proto"
)

// ITodoRepository ...
type ITodoRepository interface {
	Find(int32) (*pb.Todo, error)
	FindAll() []*pb.Todo
	Create(*pb.Todo) (*pb.Todo, error)
	Update(*pb.Todo) (*pb.Todo, error)
	Delete(*pb.Todo) error
}
