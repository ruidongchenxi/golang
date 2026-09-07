package main

import (
	"context"
	"lnhgo/grpc_error_test/proto"
	"net"
	"time"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
)

type Server struct{
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(c context.Context,request *proto.HelloRequest)(*proto.HelloReply,error){
	time.Sleep(5*time.Second)
	return &proto.HelloReply{
		Message: "hello "+request.Name,
	},nil
	//return nil,status.Errorf(codes.NotFound,"记录未找到%s",request.Name)//错误处理
}
//func (s *Server)mustEmbedUnimplementedGreeterServer(){}
func main(){
	g:=grpc.NewServer()
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