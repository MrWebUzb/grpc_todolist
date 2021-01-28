package repo

import (
	"fmt"

	pb "github.com/MrWebUzb/microservice/todo_service/service/todo/proto"
)

// SimpleRepository ...
type SimpleRepository struct {
	todos []*pb.Todo
}

// FindAll ...
func (repo *SimpleRepository) FindAll() []*pb.Todo {
	return repo.todos
}

// Find ...
func (repo *SimpleRepository) Find(id int32) (*pb.Todo, error) {
	for _, todo := range repo.todos {
		if todo.Id == id {
			return todo, nil
		}
	}

	return &pb.Todo{}, fmt.Errorf("unable to found todo with %d", id)
}

// Create ...
func (repo *SimpleRepository) Create(todo *pb.Todo) (*pb.Todo, error) {
	todo.Id = int32(len(repo.todos)) + 1
	repo.todos = append(repo.todos, todo)

	return todo, nil
}

// Update ...
func (repo *SimpleRepository) Update(todo *pb.Todo) (*pb.Todo, error) {
	old, err := repo.Find(todo.Id)
	if err != nil {
		return &pb.Todo{}, err
	}

	old = todo

	return old, nil
}

// Delete ...
func (repo *SimpleRepository) Delete(todo *pb.Todo) error {
	var updated []*pb.Todo

	for _, t := range repo.todos {
		if t.Id == todo.Id {
			continue
		}
		updated = append(updated, t)
	}

	repo.todos = updated
	return nil
}
