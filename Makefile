proto:
	protoc -I. --go_out=./ --go-grpc_out=./ service/todo/proto/todo.proto

server_build:
	go build ./cmd/server

client_build:
	go build ./cmd/client

.PHONY: proto
.DEFAULT_GOAL:=proto