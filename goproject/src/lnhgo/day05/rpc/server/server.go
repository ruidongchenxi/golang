package main

import (
	"encoding/json"
	"fmt"
	//"net"
	"net/http"
	"strconv"
)

func Add(a, b int) int {
	return a + b
}

type Company struct {
	Name    string
	Address string
}
type Employee struct {
	Name    string
	Company Company
}
type PrintResult struct{
	Info string
	Err error

}
// func RpcPrint(employee Employee) PrintResult{
// 	/*序列化（数据编码）和传输协议以及ciid
// 	http1.x http2.0协议：
// 	http协议底层使用的也是tcp协议。http协议现在主流的是1.x这种协议有性能问题，一次性一旦结果返回连接就端口，建立一次TCP连接成本比较高，http协议是在TCP连接基础之上的
// 	1. 直接基于tcp/udp 协议去封装一层协议，没有通用性，http2.0，既有http特性也有长连接特性，grpc支持http2.0
// 	客户端
// 	1.建立连接tcp/http
// 	2.将employee 对象序列化成json字符串
// 	3.发送字符串->调用成功后返回实际你接收的是二进制的数据
// 	4.等待服务返回调用结果
// 	5.将服务器返回的结果解析成PrintResult对象->反序列化
// 	服务端
// 	1.监听网络端口
// 	2.读取数据->二进制的json数据
// 	3.对数据进行反序列化成Eemployee对象
// 	4.开始处理业务逻辑
// 	5.将处理的结果PrintResult序列化成json二进制数据->序列化
// 	6.将数据返回
// 	序列号与反序列化可以选择的，不一定要采用json格式、xml、protobul、msgpack
// 	rpc中第二点
// 	网络连接（tcp/http)
// 	http协议1.0来说，有一个问题一次连接一旦对方返回了结果；连接端口
// 	http协议2.0：保持长连接grpc
// 	*/

// }
func main() {
	//将这个打印工作放在另一台服务器上，需要将本地的内存对象struct 、这样不行，可行方式将struct 序列化成json；通过二进制传输
	// fmt.Println(Employee{
	// 	Name: "go 工程师",
	// 	Company: Company{
	// 	Name: "晨曦",
	// 	Address: "北京",
	// },
	// },)
	//将远程的服务需要将二进制对象反解成struct对象
	//callID 的问题：r.URL.Path；数据的传输协议http协议，3网络传输协议tcp
	http.HandleFunc("/add",func(w http.ResponseWriter, r *http.Request) {
		_=r.ParseForm()//解析参数
		fmt.Println("path:",r.URL.Path)//获取请求路径
		a,_:=strconv.Atoi(r.Form["a"][0])//将字符串转为int
		//strconv.Itoa(a)将int转为字符串
		b,_:=strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Cont-Type","application/json")
		jData,_:=json.Marshal(map[string]int{
			"data":a+b,
		})
		_,_=w.Write(jData)


	})
	_=http.ListenAndServe(":8000",nil)
}