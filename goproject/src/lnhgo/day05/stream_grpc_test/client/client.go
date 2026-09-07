package main

import (
	"context"
	"fmt"
	"time"
	"sync"

	"lnhgo/day05/stream_grpc_test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
func main() {
	conn, err := grpc.NewClient(
		"127.0.0.1:50052",
    	grpc.WithTransportCredentials(insecure.NewCredentials()),
	)	
	if err !=nil{
		panic("连接失败")
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	res,_:=c.GetStream(context.Background(),&proto.StreamReqDate{Data: "慕课网",},)
	for{
		a,err:=res.Recv()
		if err !=nil{
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}
	//客户端流模式
	putS,_:=c.PostStream(context.Background())
	i:=0
	for{
		i++
		_=putS.Send(&proto.StreamReqDate{
			Data: fmt.Sprintf("go 中文网%d",i),

		})
		time.Sleep(time.Second)
		if i>10{
			break
		}

	}
	//双向流模式
	allStr,_:=c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data,_:=allStr.Recv()
			fmt.Println("收到客户端消息:"+data.Data)
		}
	}()
	go func ()  {
		defer wg.Done()
		for {
			_=allStr.Send(&proto.StreamReqDate{Data: "你好我是go"})
			time.Sleep(time.Second)
		}
	}()
	defer wg.Wait()

}
