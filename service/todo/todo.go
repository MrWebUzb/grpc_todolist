package todo

import (
	"context"

	pb "github.com/MrWebUzb/microservice/todo_service/service/todo/proto"
	"github.com/MrWebUzb/microservice/todo_service/service/todo/repo"
)

// Service ...
type Service struct {
	pb.UnimplementedTodoServiceServer
	repo repo.ITodoRepository
}

// New ...
func New(repo repo.ITodoRepository) *Service {
	return &Service{repo: repo}
}

// GetTodos ...
func (s *Service) GetTodos(ctx context.Context, req *pb.EmptyRequest) (*pb.Response, error) {
	return &pb.Response{
		Todos: s.repo.FindAll(),
	}, nil
}

// Get ...
func (s *Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	todo, err := s.repo.Find(req.Id)
	if err != nil {
		return &pb.Response{}, err
	}

	return &pb.Response{Todo: todo}, nil
}

// Create ...
func (s *Service) Create(ctx context.Context, req *pb.CreateUpdateRequest) (*pb.Response, error) {
	todo, err := s.repo.Create(req.Todo)

	if err != nil {
		return &pb.Response{
			Created: false,
		}, err
	}

	return &pb.Response{
		Created: true,
		Todo:    todo,
	}, nil
}

// Update ...
func (s *Service) Update(ctx context.Context, req *pb.CreateUpdateRequest) (*pb.Response, error) {
	todo, err := s.repo.Update(req.Todo)
	if err != nil {
		return &pb.Response{
			Updated: false,
		}, err
	}

	return &pb.Response{
		Updated: true,
		Todo:    todo,
	}, nil
}

// Delete ...
func (s *Service) Delete(ctx context.Context, req *pb.CreateUpdateRequest) (*pb.DeleteResponse, error) {
	err := s.repo.Delete(req.Todo)
	return &pb.DeleteResponse{Deleted: err != nil}, nil
}
