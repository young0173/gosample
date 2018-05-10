package gorpc

import (
	"helloworld"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type hello struct{}

func (h *hello) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {

	res := &helloworld.HelloResponse{
		Name: "Well.",
	}

	return res, nil

}

//Server receive message and response
func Server() {

	listener, _ := net.Listen("tcp", Address)

	server := grpc.NewServer()
	helloworld.RegisterGreeterServer(server, &hello{})
	server.Serve(listener)

	log.Println("server listen on", Address)

}
