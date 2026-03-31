package main

import (
	"fmt"
	_"io"
	"net"
)
func main(){
	fmt.Println("开始监听了")
	//1.TCP表示使用的协议是TCP的
	//2.0.0.0.0:8888 表示本机所有地址的8888端口
	lisen, err :=net.Listen("tcp","0.0.0.0:8888")
	if err !=nil{
		fmt.Println("lisen err=",err)
		return
	}
	fmt.Printf("lisen=%v\n",lisen)
	defer lisen.Close() //延迟关闭监听
	for{
		fmt.Println("等待客户端连接")
		conn, err :=lisen.Accept() //等待客户端连接

		if err != nil{
			fmt.Println("监听失败")
		}else{
			fmt.Printf("Accept() suc con=%v 客户端IP=%v\n",conn,conn.RemoteAddr().String())
		}
		//这里启动协程响应客户端请求
		go func (con net.Conn)  {
		defer con.Close()//一定关闭
		for {
			//
			buy :=make([]byte,1024)
			//con.Read(buy) 等待客户端发生信息
			//如果客户端，没有发送数据，那么协程会阻塞这里
			//fmt.Printf("服务器在等待客户端%s输入\n",con.RemoteAddr().String())
			n,err :=con.Read(buy)
			if  err !=nil{
				fmt.Println("客户端关闭")
				return
			}
			//显示客户端发送的内容到服务器终端
			fmt.Println(string(buy[:n]))
		}
		
	}(conn)
	}
	
}

//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter19\tcpdome01\server\main.go
// 开始监听了
// lisen=&{0xc000070a08 {<nil> 0 0}}
// 等待客户端连接
// Accept() suc con=&{{0xc000070c88}} 客户端IP=10.9.8.207:1597
// 等待客户端连接
// 服务器在等待客户端10.9.8.207:1597输入
// nihao

// 服务器在等待客户端10.9.8.207:1597输入
// 客户端关闭
// Accept() suc con=&{{0xc000070f08}} 客户端IP=10.9.8.207:6912
// 等待客户端连接
// 服务器在等待客户端10.9.8.207:6912输入
// utg

// 服务器在等待客户端10.9.8.207:6912输入
// 客户端关闭
