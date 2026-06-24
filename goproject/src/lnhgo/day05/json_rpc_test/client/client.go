package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//建立链接
	conn, err := net.Dial("tcp","localhost:1234")
	if err !=nil{
		panic("链接建立失败")
	}
	//var reply *string=new(string)
	var reply string
	client:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err=client.Call("HelloService.Hello","bobby",&reply)
	if err !=nil{
		panic("调用失败")
	}
	fmt.Println(reply)
	//一连串的代码大部分问题都是net的包好像跟和rpc没啥关系
	//不行。rpc调用中有几个问题需要解决1.call ID 2.序列化和反序列化编码和解码
	//python 下的开发语言这个显得不好用
	//可以跨语言调用，go语言rpc的序列化和反序列协议是什么（Gob）；2.能否替换常见的序列化协议
}