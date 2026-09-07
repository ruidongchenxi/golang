package main

import (
	"context"
	"fmt"
	"lnhgo/metadata_test/proto"
	//"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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
	//timestampFormat := ""
	//md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md := metadata.New(map[string]string{
		"name":"bobby",
		"pasword":"imooc",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//r,err:=c.SayHello(context.Background(),&proto.HelloRequest{Name: "chenxi"})
	//r,err:=c.SayHello(context.Background(),&proto.HelloRequest{})
	r,err:=c.SayHello(ctx,&proto.HelloRequest{Name: "bobby"})

	if err != nil{
		panic("调用失败")
	}
	fmt.Println(r.Message)
}