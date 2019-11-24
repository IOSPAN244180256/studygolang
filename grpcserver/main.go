package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"grpcserver/proto/test"
)

type HelloService struct {
}

// SayHello is a function
func (hs HelloService) SayHello(ctx context.Context, in *test.HelloRequest) (*test.HelloResponse, error) {
	return &test.HelloResponse{Message: fmt.Sprintf("你好，%s", in.Username)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":6021")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := grpc.NewServer()
	test.RegisterHelloServiceServer(s, HelloService{})
	s.Serve(lis)
}
