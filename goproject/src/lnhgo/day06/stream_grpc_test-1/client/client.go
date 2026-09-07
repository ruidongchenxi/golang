package main

import (
	"context"
	"fmt"
	"time"

	"lnhgo/day06/stream_grpc_test-1/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
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
	//r,err:=c.SayHello(context.Background(),&proto.HelloRequest{Name: "chenxi"})
	r,err:=c.SayHello(context.Background(),&proto.HelloRequest{
		Name: "bobby",
		Url:"vvv",
		G: proto.Gender_FEMALE,
		Mp: map[string]string{
			"name":"boov",
			"coon":"vvv",
		},
		AddTime: timestamp.New(time.Now()),

	})
	if err != nil{
		panic("调用失败")
	}
	fmt.Println(r.Name)
}