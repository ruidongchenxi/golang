package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"src/chatroom/common/message"
)
func ReadPkg(conn net.Conn) (mes message.Message,err error){
	buf:=make([]byte,8096)
	fmt.Println("等待客户端读取数据")
	//conn.Read 在con没有被关闭情况下，才会阻塞
	//如果有任意一方关闭连接，则不会阻塞，返回
	_,err=conn.Read(buf[:4])
	if err!=nil{
		//err = errors.New("red pkg header error")
		return
	}
	//fmt.Println("读到的内容=",buf[:4])
	//根据读到buf[:4] 转成uint32类型
	pkgLen :=binary.BigEndian.Uint32(buf)
	//根据pkgLen 读取消息内容
	n,err := conn.Read(buf[:pkgLen])//从连接里读数据写到buf 切片里
	if n!=int(pkgLen)|| err!=nil{
		err = errors.New("red pkg body error")
		return 
	}
	//把buf切片反序列化
	//var mes message.Message
	//技术就是一层窗户纸
	err=json.Unmarshal(buf[:pkgLen],&mes)//注意是指针这里
	if err !=nil{
		fmt.Println("反序列化失败",err)
		return
	}
	return

}
func Process(conn net.Conn)  {
	//这里需要延迟关闭连接
	defer conn.Close()
	// //读取客户端发送信息
	// 
	for {
	//这里将读取数据包，直接封装成一个函数readPkg(),返回Messahe,err
	//等待读取
		mes,err:= ReadPkg(conn)
		if err !=nil {
			if err == io.EOF{
				fmt.Println("客户端关闭连结，服务端也关闭")
				return
			}
			fmt.Println("获取连接序列化失败",err)
			return
		}
		fmt.Println("mes=",mes)

	}

}
func main(){
	//
	fmt.Println("服务器监听8889端口")
	listen,err:=net.Listen("tcp","0.0.0.0:8889")
	
	if err != nil{
		fmt.Println("服务器端监听失败")
		return
	}
	defer listen.Close()
	//一旦监听成功等待客户端连接
	for{
		fmt.Println("等待客户端连接")
		conn,err:=listen.Accept()
		if err != nil{
			fmt.Println("连接失败")
		}
		//一旦连接成功，就启动协程和客户端保持通讯
		//处理通讯信息
		go Process(conn)
	}
}