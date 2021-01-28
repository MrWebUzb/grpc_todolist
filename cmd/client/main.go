package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	pb "github.com/MrWebUzb/microservice/todo_service/service/todo/proto"
	"google.golang.org/grpc"
)

const (
	defaultFile = "todos.json"
	defaultPort = 9999
	defaultHost = "localhost"
)

func main() {
	var filename string
	var host string
	var port int
	var create, getTodos bool

	flag.StringVar(&filename, "file", defaultFile, "todo list json file for adding to server")
	flag.StringVar(&host, "host", defaultHost, "server host")
	flag.IntVar(&port, "port", defaultPort, "server port")
	flag.BoolVar(&create, "create", false, "create todo list from given file")
	flag.BoolVar(&getTodos, "get-all", false, "get all todo lists from server")

	flag.Parse()

	conn, err := getConnection(host, port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create connection %s: %v\n", fmt.Sprintf("%s:%d", host, port), err)
		os.Exit(1)
	}

	client := pb.NewTodoServiceClient(conn)

	if create {
		_, err := createFromFile(client, filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to create todos: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "successfully created!")
	}

	if getTodos {
		todos, err := getAll(client)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to create todos: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("here is your todo list:\n")
		for _, todo := range todos {
			fmt.Printf("%d: %v - %v [%v]\n", todo.Id, todo.Title, todo.Description, todo.Done)
		}
	}
}

func getAll(client pb.TodoServiceClient) ([]*pb.Todo, error) {
	res, err := client.GetTodos(context.Background(), &pb.EmptyRequest{})

	if err != nil {
		return []*pb.Todo{}, err
	}

	return res.Todos, nil
}

func createFromFile(client pb.TodoServiceClient, filename string) (bool, error) {
	handler, err := openFile(filename)
	if err != nil {
		return false, err
	}
	defer handler.Close()

	todos, err := parseFile(handler)
	if err != nil {
		return false, err
	}

	result := true

	for _, todo := range todos {
		res, err := client.Create(context.Background(), &pb.CreateUpdateRequest{
			Todo: todo,
		})
		if err != nil {
			return false, err
		}

		result = result && res.Created
	}

	return result, nil
}

func getConnection(host string, port int) (grpc.ClientConnInterface, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())

	return conn, err
}

func openFile(filename string) (*os.File, error) {
	handler, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return handler, nil
}

func parseFile(file *os.File) ([]*pb.Todo, error) {
	data, err := ioutil.ReadAll(file)

	if err != nil {
		return []*pb.Todo{}, err
	}
	todos, err := parseData(data)
	if err != nil {
		return []*pb.Todo{}, err
	}

	return todos, nil
}

func parseData(data []byte) ([]*pb.Todo, error) {
	var todos []*pb.Todo

	err := json.Unmarshal(data, &todos)

	return todos, err
}
