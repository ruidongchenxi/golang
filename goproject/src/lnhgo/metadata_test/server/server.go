package main

import (
	"context"
	"fmt"
	"lnhgo/metadata_test/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct{
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(c context.Context,request *proto.HelloRequest)(*proto.HelloReply,error){
	//matadata := nil
	md,ok:=metadata.FromIncomingContext(c)
	if ok {
		fmt.Println("get metadata error")
	}
	for key,val:=range md{
		fmt.Println(key,val)
	}
	if name,ok:=md["name"];ok{
		fmt.Println(name)
	}
	return &proto.HelloReply{
		Message: "hello "+request.Name,
	},nil
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