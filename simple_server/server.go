package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"hexmeet.com/grpctest/proto"
)

type HelloService struct {
}

func (s *HelloService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {

	md, ok := metadata.FromIncomingContext(ctx)

	fmt.Println(ok)
	if ok {
		for key, val := range md {
			fmt.Printf("%v, %v \n", key, val)
		}
	}

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

	fmt.Println(reply)
	//return reply, nil

	return nil, status.Error(codes.InvalidArgument, "invalid argument")

}

func main() {

	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//对req中的签名进行验证
		//如果失败直接返回Nil，err=验证签名失败

		//handler是客户端原来打算调用的方法，如果验证成功，执行真正的方法
		fmt.Println("recving a request")
		res, err := handler(ctx, req)
		fmt.Println("done")
		return res, err
	}

	var opts []grpc.ServerOption //grpc为使用的第三方的grpc包
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	//opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opts...)
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
