package main

import (
	"context"
	"fmt"
	"lnhgo/metadata_test/proto"
	"time"

	//"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	interceptor:=func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
		fmt.Println("==========进入客户端拦截器==========")
    	
		fmt.Println("method:", method)
		start:= time.Now()
		err:=invoker(ctx,method,req,reply,cc,opts...)
		fmt.Printf("耗时：%s\n",time.Since(start))
		return err

	}
	ops:=grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.NewClient(
    "127.0.0.1:8080",
    grpc.WithTransportCredentials(insecure.NewCredentials()),
	ops,
)
	if err !=nil{
		panic("连接失败")
	}
	defer conn.Close()
	c:=proto.NewGreeterClient(conn)
	md := metadata.New(map[string]string{
		"name":"bobby",
		"pasword":"imooc",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r,err:=c.SayHello(ctx,&proto.HelloRequest{Name: "bobby"})

	if err != nil{
		panic("调用失败")
	}
	fmt.Println(r.Message)
}