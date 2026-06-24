package main

import (
	//"net"
	"net"
	"net/rpc"
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
	conn,_:=listener.Accept()//当一个新链接进来的时候
	rpc.ServeConn(conn)
}