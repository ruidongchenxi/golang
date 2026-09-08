package main

import (
	"flag"
	"fmt"
	"mxshop_srvs/user_srv/handler"
	"mxshop_srvs/user_srv/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	IP:=flag.String("ip","0.0.0.0","ip地址")//采用这种方法可以灵活指定IP
	Prot:=flag.Int("prot",50051,"端口：")//采用这种方法可以灵活指定端口
	flag.Parse()
	fmt.Println("IP",*IP)
	fmt.Println("prot",*Prot)

	server := grpc.NewServer()
	proto.RegisterUserServer(server,&handler.UserServer{})
	list,err:= net.Listen("tcp",fmt.Sprintf("%s:%d",*IP,*Prot))//设置监听
	if err!=nil{
		panic("failed to listen"+err.Error())
	}
	err = server.Serve(list)
	if err != nil{
		panic("faile to start grpc:"+err.Error())
	}
}