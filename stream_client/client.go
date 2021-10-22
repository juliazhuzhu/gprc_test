package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"hexmeet.com/grpctest/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := proto.NewJreeterClient(conn)
	//Server Stream mode
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "ffu"})
	for {
		data, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(data.Data)
	}

	//Client Stream mode
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("imoc %d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//Bi Stream mode
	biStr, _ := c.AllStream(context.Background())
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
			biStr.Send(&proto.StreamReqData{
				Data: "hello all stream client.",
			})

			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

}
