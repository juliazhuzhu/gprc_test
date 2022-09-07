package main

import (
	"context"
	"fmt"
	"hexmeet.com/grpctest/grpclb_test/proto"
	"log"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"consul://172.20.0.204:8500/user-srv?wait=14s&tag=imoc",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i:=0; i < 10; i++ {

		userSrvClient := proto.NewUserClient(conn)
		rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
			Pn:    1,
			PSize: 2,
		})

		if err != nil {
			panic(err)
		}

		for index, data := range rsp.Data {
			fmt.Println(index, data)
		}
	}
}
