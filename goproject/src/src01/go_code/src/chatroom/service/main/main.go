package main

import (
	//"encoding/binary"
	_"encoding/json"
	//"errors"
	"fmt"
	_"io"
	"net"
	_"src/chatroom/common/message"
	_"src/chatroom/service/utils"
)
// func ReadPkg(conn net.Conn) (mes message.Message,err error){
// 	buf:=make([]byte,8096)
// 	fmt.Println("等待读取数据~~~~")
// 	//conn.Read 在con没有被关闭情况下，才会阻塞
// 	//如果有任意一方关闭连接，则不会阻塞，返回
// 	_,err=conn.Read(buf[:4])
// 	if err!=nil{
// 		//err = errors.New("red pkg header error")
// 		return
// 	}
// 	//fmt.Println("读到的内容=",buf[:4])
// 	//根据读到buf[:4] 转成uint32类型
// 	pkgLen :=binary.BigEndian.Uint32(buf)
// 	//根据pkgLen 读取消息内容
// 	n,err := conn.Read(buf[:pkgLen])//从连接里读数据写到buf 切片里
// 	if n!=int(pkgLen)|| err!=nil{
// 		err = errors.New("red pkg body error")
// 		return 
// 	}
// 	//把buf切片反序列化
// 	//var mes message.Message
// 	//技术就是一层窗户纸
// 	err=json.Unmarshal(buf[:pkgLen],&mes)//注意是指针这里
// 	if err !=nil{
// 		fmt.Println("反序列化失败",err)
// 		return
// 	}
// 	return

// }
// func WritePkg(conn net.Conn,data []byte) (err error){
// 	//发送长度给对方
// 	pkgLen:=uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
// 	n,err:=conn.Write(buf[:4])
// 	if  n!=4 || err!=nil {
// 		fmt.Println("发生失败")
// 		return err
// 	}
// 	//发送data 本身
// 	n,err=conn.Write(data)
// 	if n!=int(pkgLen) || err != nil{
// 		fmt.Println("发送失败",err)
// 		return
// 	}

// 	return

// }
//编写一个函数servicePro,专门处理登录请求
// func serviceProcessLogin(coon net.Conn,mes *message.Message)(err error){
// 	//核心代码
// 	//先从mes中取出mes.data,并反序列化loginmes
// 	var loginMes message.LoginMes 
// 	err=json.Unmarshal([]byte(mes.Data),&loginMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal fail err=",err)
// 		return 
// 	}
// 	//先声明一个返回的结构体
// 	var resMes message.Message
// 	resMes.Type= message.LoginResMesType
// 	//声明一个LoginResM
// 	var  loginResMes message.LoginResMes
// 	//如果用户id=100 密码=123456；认为合法，否则不合法
// 	if loginMes.UserId==100&& loginMes.UserPwd=="123456"{
// 		fmt.Println("用户合法")
// 		loginResMes.Code=200//合法
// 	}else{
// 		//fmt.Println("用户不合法")
// 		loginResMes.Code =500 //不合法
// 		fmt.Println("非法用户")
// 		loginResMes.Error = "非法用户"

// 	}
// 	//loginResMes 序列化
// 	data,err:=json.Marshal(loginResMes)
// 	if err !=nil {
// 		fmt.Println("序列化失败",err)
// 		return
// 	}
// 	//将data 赋值给
// 	//将data 赋值给
// 	resMes.Data = string(data)
// 	//对他序列化
// 	data,err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("返回的数据序列化失败",err)
// 	}
// 	//发送data
// 	err = WritePkg(coon,data)
// 	return
// }
//编写一个serviceProcessMes函数，功能根据不同类型消息调不同函数处理
// func serviceProcessMes(coon net.Conn,mes *message.Message) (err error){
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		//处理登录函数
// 		err =serviceProcessLogin(coon,mes)
// 	case message.RegisterMesType:
// 		//处理注册
// 	default:
// 		fmt.Println("消息类型不存在，无法处理")

// 	}
// 	return 
// }
func Process(conn net.Conn)  {
	//这里需要延迟关闭连接
	defer conn.Close()
	// //读取客户端发送信息
	// 
		
	// for {
	// //这里将读取数据包，直接封装成一个函数readPkg(),返回Messahe,err
	// //等待读取
	// 	mes,err:= utils.ReadPkg(conn)
	// 	if err !=nil {
	// 		if err == io.EOF{
	// 			fmt.Println("客户端关闭连结，服务端也关闭")
	// 			return
	// 		}else{
	// 			fmt.Println("获取连接序列化失败",err)
	// 			return
	// 		}
	// 	}
	// //	fmt.Println("mes=",mes)
	// 	err = serviceProcessMes(conn,&mes)
	// 	if err !=nil {
	// 		return
		
	// 	}
	// }
	processor:= &Processor{
		Conn: conn,
	}
	err :=processor.Process2()
	if err !=nil {
		fmt.Println("客户端通讯协议错误=err",err)
	}
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