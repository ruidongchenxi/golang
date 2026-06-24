package main

import (
	//"net"
	//"encoding/json"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)
type HelloService struct{

}
func (s *HelloService)Hello(requst string,reply *string) error{
	*reply = "hello,"+requst
	return nil
}

func main() {
	//1.实例化一个server
	listener,_:=net.Listen("tcp",":1234")
	//2.注册处理逻辑
	
	_=rpc.RegisterName("HelloService",&HelloService{})  //注册rpc
	//启动服务
	for{
		conn,_:=listener.Accept()//当一个新链接进来的时候
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}