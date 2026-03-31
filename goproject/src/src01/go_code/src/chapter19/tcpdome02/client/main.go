package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)
func main(){
	conn,err := net.Dial("tcp","10.9.8.207:8888")
	if err != nil{
		fmt.Println("连接失败err=",err)
		return
	}
	fmt.Println("conn成功=",conn)
	//准备和客户端发生单行数据
	reader :=bufio.NewReader(os.Stdin)
	//从终端读取用户的输入，并发送到服务器断
	str,err :=reader.ReadString('\n')
	if err !=nil{
		fmt.Println("readString err",err)
	}
	//将str 发送给服务器
	n,err :=conn.Write([]byte(str))//将str 转成切片后发给服务器
	if err !=nil{
		fmt.Printf("发送信息失败err=\n",err)
	}
	fmt.Printf("客户端发生了=%v字节,并退出\n",n)
}



// PS D:\golang\goproject\src\src01\go_code\src> go  run chapter19\tcpdome01\client\main.go
// conn成功= &{{0xc000070a08}}
// nihao
// 客户端发生了=7字节,并退出
// PS D:\golang\goproject\src\src01\go_code\src> go  run chapter19\tcpdome01\client\main.go
// conn成功= &{{0xc000070a08}}
// utg
// 客户端发生了=5字节,并退出
