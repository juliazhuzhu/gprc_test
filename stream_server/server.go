package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"hexmeet.com/grpctest/proto"
)

const PORT = ":50052"

type server struct {
}

func (s *server) GetStream(req *proto.StreamReqData, res proto.Jreeter_GetStreamServer) error {

	i := 0
	for {
		i++
		fmt.Println(req.Data)
		res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})

		time.Sleep(time.Second)
		if i > 10 {
			break
		}

	}

	return nil
}

func (s *server) PutStream(cliStr proto.Jreeter_PutStreamServer) error {

	for {
		if data, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(data.Data)
		}

	}
	return nil
}

func (s *server) AllStream(biStr proto.Jreeter_AllStreamServer) error {
	//return nil
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			if data, err := biStr.Recv(); err != nil {
				fmt.Println(err)
				break
			} else {
				fmt.Println(data.Data)
			}

		}
	}()

	go func() {
		defer wg.Done()
		for {
			biStr.Send(&proto.StreamResData{
				Data: "hello all stream server.",
			})

			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

	return nil
}

func main() {

	g := grpc.NewServer()
	proto.RegisterJreeterServer(g, &server{})
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	err = g.Serve(lis)
	if err != nil {
		panic("failed to start gprc:" + err.Error())
	}
}
