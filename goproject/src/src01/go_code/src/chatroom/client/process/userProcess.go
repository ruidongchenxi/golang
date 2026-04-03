package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"src/chatroom/client/utils"
	"src/chatroom/common/message"
)
type UserProcess struct{
}

func (this *UserProcess)Register(userId int,
	userPwd string,userName string) (err error){
	conn,err:=net.Dial("tcp","localhost:8889")
	if err != nil{
		fmt.Println("连接失败")
		return err
	}
	//延迟关闭
	defer conn.Close()
	var mes message.Message
	mes.Type=message.RegisterResMesType
	var RegisterMes message.RegisterMes
	// RegisterResMes.User.UserId=userId
	// RegisterResMes.User.UserPwd=userPwd
	// RegisterResMes.User.UserName=userName
	RegisterMes.User.UserID=userId
	RegisterMes.User.UserPwd=userPwd
	RegisterMes.User.UserName=userName
	data,err:=json.Marshal(RegisterMes)
	if err !=nil{
		fmt.Println("序列化失败")
		return		
	}
	//5.把data 赋给了mes.Data
	mes.Data=string(data)
	data,err= json.Marshal(mes)
	if err != nil{
		fmt.Println("消息序列化失败")
		return 
	}
    tf := &utils.Transfer{
		Conn: conn,
	}
	pkgLen:=uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	n,err:=conn.Write(buf[:4])
	if  n!=4 || err!=nil {
		fmt.Println("发生失败")
		return err
	}
	 fmt.Println("客户端发送长度:",len(data))
	 fmt.Println("客户端发内容:",string(data))
	//发生
	err=tf.WritePkg(data)
	if err !=nil{
		fmt.Println("注册失败",err)
	}
	mes,err=tf.ReadPkg()//
	if err !=nil{
		//接收失败
		fmt.Println("接收失败",err)
		return
	}
	var RegisterResMes  message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&RegisterResMes)
	//json.Unmarshal([]byte(mes.Data),&loginResMes)
	// if loginResMes.Code==200 {
	// 	fmt.Println("登录成功")
	// }else if loginResMes.Code==500{
	// 	fmt.Println("失败",loginResMes.Error)
	// }
	switch RegisterResMes.Code{
	case 200:
		//启动一个协程，该协程保持和服务器端的通信
		//该协程保持和服务器端的通讯，如果服务器有数据推送；给客户端可以接受并显示客户端终端
		// fmt.Println("登录成功") 显示登录成功后菜单循环
		fmt.Println("注册成功")
		os.Exit(0)
	default:
		fmt.Println("失败",RegisterResMes.Error)
		os.Exit(0)
	}
	
	return
}

func (this *UserProcess)Login(id int,pwd string) (err error){
   //下一步定义协议
//    fmt.Println("你好")
//    return nil
// 连接服务器
	conn,err:=net.Dial("tcp","localhost:8889")
	if err != nil{
		fmt.Println("连接失败")
		return err
	}
	//延迟关闭
	defer conn.Close()
	//准备通过con 发生消息
	var mes message.Message
	mes.Type=message.LoginMesType
	//创建一个loginmes
	var loginMes message.LoginMes
	loginMes.UserId=id
	loginMes.UserPwd=pwd
	data,err:=json.Marshal(loginMes)
	if err !=nil{
		fmt.Println("序列化失败")
		return err		
	}
	//5.把data 赋给了mes.Data
	mes.Data=string(data)
	//6.将mes序列化
	data,err= json.Marshal(mes)
	if err != nil{
		fmt.Println("消息序列化失败")
		return err
	}
	//data 就是是我们要发生的数据；根据原先规则先发送一个长度
	//先把data 长度，将data长度转成可以表示长度byte切片
	pkgLen:=uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	n,err:=conn.Write(buf[:4])
	if  n!=4 || err!=nil {
		fmt.Println("发生失败")
		return err
	}
	// fmt.Println("消息长度发生成功")
	 fmt.Println("客户端发送长度:",len(data))
	 fmt.Println("客户端发内容:",string(data))
	//发送消息本身
	n,err=conn.Write(data)
	if err !=nil {
		fmt.Println("发送失败")
		return
	}
	//这里需要处理服务器端返回的消息
	//创建结构体示例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes,err=tf.ReadPkg()//
	if err !=nil{
		//接收失败
		fmt.Println("接收失败",err)
		return
	}
	//将mes 的data 部分反序列化
	var  loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)


	switch loginResMes.Code{
	case 200:
		//启动一个协程，该协程保持和服务器端的通信
		//该协程保持和服务器端的通讯，如果服务器有数据推送；给客户端可以接受并显示客户端终端
		// fmt.Println("登录成功") 显示登录成功后菜单循环
		go ServerProcessMes(conn)

		for {
			ShowMenu()
		}
	default:
		fmt.Println("失败",loginResMes.Error)
	}
	return 
}