package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hexmeet.com/grpctest/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	/*reply, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name:    "xiaoye",
		GenTime: timestamppb.Now(),
		Gender:  proto.Gender_MALE,
		Infomation: &proto.HelloRequest_Info{
			Info: "you are the man.",
		},
		Dune: map[string]string{
			"name": "Rebeca",
		},
		Cast: []int32{2, 3},
	})*/

	hello := proto.HelloRequest{
		Name:    "xiaoye",
		GenTime: timestamppb.Now(),
		Gender:  proto.Gender_MALE,
		Infomation: &proto.HelloRequest_Info{
			Info: "you are the man.",
		},
		//Dune: map[string]string{
		//	"name": "Rebeca",
		//},
		//Cast: []int32{2, 3},
	}
	hello.Dune = make(map[string]string)
	hello.Dune["name"] = "yulian"

	//hello.Cast = []int32{}
	hello.Cast = append(hello.Cast, 9)

	reply, err := c.SayHello(context.Background(), &hello)
	if err != nil {
		panic(err)
	}
	fmt.Println((reply.Message))
	fmt.Println("end")
}
