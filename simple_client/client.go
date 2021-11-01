package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"hexmeet.com/grpctest/proto"
)

type customRPCCredentials struct {
}

func (c *customRPCCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {

	return map[string]string{
		"appid":  "100010",
		"appkey": "1adkfafadf",
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (c *customRPCCredentials) RequireTransportSecurity() bool {
	return false
}

func main() {

	/*interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println("consumed: ,", time.Since(start))
		return err
	}

	opt := grpc.WithUnaryInterceptor(interceptor)*/
	opt := grpc.WithPerRPCCredentials(&customRPCCredentials{})
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure(), opt)
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

	//md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))

	md := metadata.New(map[string]string{
		"token": "232342323kjj;",
	})

	timeout_ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()
	ctx := metadata.NewOutgoingContext(timeout_ctx, md)

	reply, err := c.SayHello(ctx, &hello)
	if err != nil {
		//panic(err)
		st, ok := status.FromError(err)
		if !ok {
			panic("parse err failure")
		}

		fmt.Println(st.Message())
		fmt.Println(st.Code())

		return

	}
	fmt.Println((reply.Message))
	fmt.Println("end")
}
