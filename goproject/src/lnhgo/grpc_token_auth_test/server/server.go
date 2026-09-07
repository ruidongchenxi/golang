package main

import (
	"context"
	"fmt"
	"lnhgo/grpc_token_auth_test/proto"
	"net"

	//"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"//如果没有status 请执行下载命令go get google.golang.org/grpc 和 go get google.golang.org/grpc/status
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
		md,ok:=metadata.FromIncomingContext(ctx)
		fmt.Println(md)
		if !ok {
			return resp,status.Error(codes.Unauthenticated,"无token认证信息")
			//fmt.Println("get metadata error")
		}
		var (
			appid string
			appkey string
		)
		if va1,ok:=md["appid"];ok {
			appid = va1[0]
			fmt.Println(appid)
		}
		if va2,ok:=md["appkey"];ok {
			appkey = va2[0]
			fmt.Println(appkey)
		}
		if appid != "101010"|| appkey != "i am key"{
			fmt.Println("验证失败")
			return resp,status.Error(codes.Unauthenticated,"认证信息错误")
		}
		// if nameSlice,ok:=md["appid"];ok{
		// 	fmt.Println(nameSlice)
		// 	for i,e:=range nameSlice{
		// 		fmt.Println(i,e)
		// 	}
		// }
		res,err:= handler(ctx,req)
		fmt.Println("请求完成")
		return res ,err 
	}
	//grpc.WithPerRPCCredentials(customCredential{})
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