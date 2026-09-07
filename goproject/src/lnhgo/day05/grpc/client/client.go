package main

import (
	"context"
	"fmt"
	"lnhgo/day05/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
    "127.0.0.1:8080",
    grpc.WithTransportCredentials(insecure.NewCredentials()),
)
	if err !=nil{
		panic("连接失败")
	}
	defer conn.Close()
	c:=proto.NewGreeterClient(conn)
	r,err:=c.SayHello(context.Background(),&proto.HelloRequest{Name: "chenxi"})
	if err != nil{
		panic("调用失败")
	}
	fmt.Println(r.Message)
}