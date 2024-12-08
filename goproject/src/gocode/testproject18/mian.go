package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
)
func main(){
	//dayin
	fmt.Println("客户端启动")
	//调用Dial 函数;参数需要指定协议 指定服务器端的IP+PORT
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		//链接失败
		fmt.Println("错误端链接失败： err",err)
		return
	}
	fmt.Println("链接成功",conn)
	reader := bufio.NewReader(os.Stdin)//终端标准输入
    //读取用户输入
	str,err := reader.ReadString('\n')
	if err != nil {
       fmt.Println("终端输入失败,err",err)
	}
	//将终端输入数据发生给服务器
    n,err := conn.Write([]byte(str))
	if err != nil{
		fmt.Println("失败",err)
	}
	fmt.Printf("发送成功发生了%d字节\n",n)
}