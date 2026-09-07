package main

import (
	"context"
	"fmt"
	"lnhgo/grpc_interpretor/proto"
	"net"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/metadata"
)

type Server struct{
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(c context.Context,request *proto.HelloRequest)(*proto.HelloReply,error){
	//matadata := nil
	// md,ok:=metadata.FromIncomingContext(c)
	// if ok {
	// 	fmt.Println("get metadata error")
	// }
	// for key,val:=range md{
	// 	fmt.Println(key,val)
	// }
	// if name,ok:=md["name"];ok{
	// 	fmt.Println(name)
	// }
	return &proto.HelloReply{
		Message: "hello "+request.Name,
	},nil
}
//func (s *Server)mustEmbedUnimplementedGreeterServer(){}

func main(){
	interceptor:=func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error){
		fmt.Println("接收到请求")
		res,err:= handler(ctx,req)
		fmt.Println("请求完成")
		return res ,err 
	}
	opt:= grpc.UnaryInterceptor(interceptor)//server 拦截器
	g:=grpc.NewServer(opt)
	proto.RegisterGreeterServer(g,&Server{})
	lis,err:=net.Listen("tcp","0.0.0.0:8080")
	if err !=nil {
		panic("监听失败")
	}
	err=g.Serve(lis)
	if err !=nil {
		panic("启动失败")
	}
}