package main
import (	
	"src/chatroom/service/model"
	"fmt"
	"net"
)
func Process(conn net.Conn)  {
	//这里需要延迟关闭连接
	defer conn.Close()
	processor:= &Processor{
		Conn: conn,
	}
	err :=processor.Process()
	if err !=nil {
		fmt.Println("客户端通讯协议错误=err",err)
	}
}
func init(){
	initPool("127.0.0.1:6379",10,3,300)
	initUserDao()
}
//编写一个函数完成对UserDao 初始化任务
func initUserDao(){
	//这里pool 是全局变量
	//这里有一个初始化顺序问题
	//initPool，在initUserDao
	model.MyUserDao= model.NewUserDao(pool)
}
//根据用户传来的
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