package main

import (
	"context"
	"fmt"
	"lnhgo/grpc_token_auth_test/proto"
	"time"

	//"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)
type customCredential struct{}
func (c customCredential)GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error){
	return map[string]string{
			"appid":"101010",
			"appkey":"i am key",
	},nil
}
func (c customCredential)RequireTransportSecurity() bool{
	return false
}

func main() {
	interceptor:=func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
		fmt.Println("==========进入客户端拦截器==========")
    	
		fmt.Println("method:", method)
		start:= time.Now()
		md:= metadata.New(map[string]string{
			"appid":"101010",
			"appkey":"i am key",
		})
		ctx=metadata.NewOutgoingContext(context.Background(),md)
		err:=invoker(ctx,method,req,reply,cc,opts...)
		fmt.Printf("耗时：%s\n",time.Since(start))
		return err

	}
	
	ops:=grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.NewClient(
    "127.0.0.1:8080",
    grpc.WithTransportCredentials(insecure.NewCredentials()),
	ops,
	grpc.WithPerRPCCredentials(customCredential{}),
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