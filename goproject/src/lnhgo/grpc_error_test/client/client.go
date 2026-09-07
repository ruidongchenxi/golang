package main

import (
	"context"
	"fmt"
	"lnhgo/grpc_error_test/proto"
	"time"

	//"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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
	ctx,cc:= context.WithTimeout(context.Background(),time.Second*3)
	defer cc()
	_,err=c.SayHello(ctx,&proto.HelloRequest{Name: "chenxi"})
	if err != nil{
		st, ok := status.FromError(err)
		if !ok {
    // Error was not a status error
			panic("解析error 错误")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
		// panic(err)
	}
	//fmt.Println(r.Message)
}