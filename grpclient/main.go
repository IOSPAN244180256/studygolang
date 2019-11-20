package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"grpclient/proto/test"
)

func main() {
	conn, err := grpc.Dial("localhost:6021", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	c := test.NewHelloServiceClient(conn)
	r, err := c.SayHello(context.Background(), &test.HelloRequest{Username: "ft"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r.Message)
}
