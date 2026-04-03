package process

import (
	"fmt"
	"os"
	"src/chatroom/client/utils"
	"net"
)
func ShowMenu(){
	fmt.Println("---------------恭喜XXX登录成功----------------------")
	fmt.Println("---------------1.在线用户列表----------------------")
	fmt.Println("---------------2.发生消息----------------------")
	fmt.Println("---------------3.信息列表----------------------")
	fmt.Println("---------------4.退出系统----------------------")
	fmt.Println("请选择")
	var key int
	fmt.Scan(&key)
	switch key{
	case 1:
		fmt.Println("用户列表")
	case 2:
		fmt.Println("发生消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入不正确")
		//os.Exit(0)
	}
}
func ServerProcessMes(conn net.Conn){
	//创建一个结构化实例
	//
	tf :=&utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端不停地读取")
		mes,err:=tf.ReadPkg()
		if err !=nil{
			fmt.Println("读取出错")
			return
		}
		fmt.Println("mes=",mes)
	}

}