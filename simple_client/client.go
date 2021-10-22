package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"hexmeet.com/grpctest/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	reply, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "xiaoye"})
	if err != nil {
		panic(err)
	}
	fmt.Println((reply.Message))
	fmt.Println("end")
}
