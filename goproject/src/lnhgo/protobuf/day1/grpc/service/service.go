package main

import (
	"context"
	"lnhgo/protobuf/day1/grpc/proto"
	"net"

	"google.golang.org/grpc"
)
type Server struct{
	proto.UnimplementedGreeterServer
}
func (s *Server)SayHello(ctx context.Context,request *proto.HelloRequest)(*proto.HelloReply,error){
	return &proto.HelloReply{
		Message: "hello"+request.Name,
	},nil
}
func main(){
	g := grpc.NewServer()
	//注册
	proto.RegisterGreeterServer(g,&Server{})
	lis,err :=net.Listen("tcp","0.0.0.0:8088")
	if err !=nil{
		panic("端口启动失败"+err.Error())
	}
	err = g.Serve(lis)
	if err !=nil{
		panic("grpc 启动失败"+err.Error())
	}
}