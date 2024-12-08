package main
import (
	"fmt"
	"net"

)
func process(conn net.Conn){
	defer conn.Close()
	for {
		//创建切片
		buf := make([]byte,1024)
		n,err := conn.Read(buf)
		if err != nil {
           return
		}
		fmt.Println(string(buf[0:n]))
	}
}
func main(){
	//打印信息
	fmt.Println("服务器端启动")
	//进行监听
	listen,err  := net.Listen("tcp","127.0.0.1:8888")
    if err != nil{
		fmt.Println("监听失败")
	    return
	}
	//监听成功
	//等待客户端链接
	for {
		conn,err2 := listen.Accept()
		if err2 != nil{
			fmt.Println("连接失败")
		}else{
			//链接成功
			fmt.Printf("成功 IP=%v. 接受客户端信息：%v\n",conn,conn.RemoteAddr().String())
		}
		//准备协程
		go process(conn)
	}
}