package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/MrWebUzb/microservice/todo_service/service/todo"
	pb "github.com/MrWebUzb/microservice/todo_service/service/todo/proto"
	"github.com/MrWebUzb/microservice/todo_service/service/todo/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 9999, "listen port of grpcService")

	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("unable to start service: %v\n", err)
	}

	server := grpc.NewServer()
	pb.RegisterTodoServiceServer(
		server,
		todo.New(&repo.SimpleRepository{}),
	)
	reflection.Register(server)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("error serving on port :%d => :%v\n", port, err)
	}
}
