package main

import (
	//"container/list"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply 的值
	*reply = "Hello," + request
	return nil
}
func main() {
	//实例化server
	listener,_:= net.Listen("tcp",":1234")
	//注册逻辑
	_=rpc.RegisterName("HelloService",&HelloService{})
	//启动服务
	conn,_:=listener.Accept()//当新连接进来后
	rpc.ServeConn(conn)
//一串的代码net的包好像和rpc没有关系
//不行，rpc调用中有几个问题需要解决，1.call id 2.序列化和反序列化
/*
python 下的开发而言这个显得不好用
可以跨语言调用
1.go语言的rpc的序列化和反序列化是什么（Gob）
2.能否替换成常见的序列号

*/
}
