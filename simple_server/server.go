package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"hexmeet.com/grpctest/proto"
)

type HelloService struct {
}

func (s *HelloService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {

	fmt.Println(req.Name)
	fmt.Println(req.GenTime)
	fmt.Println(req.Infomation.Info)
	fmt.Println(req.Gender)
	fmt.Println(req.Dune["name"])
	for _, k := range req.Cast {
		fmt.Println(k)
	}
	reply := &proto.HelloReply{
		Message: "hahah:" + req.Name,
	}

	return reply, nil

}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &HelloService{})
	lis, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		panic(err)
	}

	err = g.Serve(lis)
	if err != nil {
		panic("failed to start gprc:" + err.Error())
	}
}
