package process2

import (
	"encoding/json"
	"fmt"
	"net"
	"src/chatroom/common/message"
	"src/chatroom/service/model"
	"src/chatroom/service/utils"
)

type UserProcess struct{
	Conn net.Conn

} 
func (this *UserProcess)ServiceProcessLogin(mes *message.Message)(err error){
	//核心代码
	//先从mes中取出mes.data,并反序列化loginmes
	var loginMes message.LoginMes 
	err=json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=",err)
		return 
	}
	//先声明一个返回的结构体
	var resMes message.Message
	resMes.Type= message.LoginResMesType
	//声明一个LoginResM
	var  loginResMes message.LoginResMes
	//需要到redis 数据库获取数据
	//使用model.MyUserDao到re
	user,err :=model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	
	if  err !=nil {
		// if err == model.ERROR_USER_NOTEXISTS{
			
		// } 
		 
		// //loginResMes.Error = "用户不存在请注册"
		switch err{
		case model.ERROR_USER_NOTEXISTS:
			loginResMes.Code =500
			loginResMes.Error = err.Error()
		case model.ERROR_USER_EXISTS:
			loginResMes.Code =403
			loginResMes.Error= err.Error()
		case model.ERROR_USER_PWD:
			loginResMes.Code =404
			loginResMes.Error= err.Error()
		default:
			loginResMes.Code =505
			loginResMes.Error="未知错误"

		}
		}else{
		fmt.Println(user,"登录成功")
		loginResMes.Code=200//合法
	}

	//如果用户id=100 密码=123456；认为合法，否则不合法
	// if loginMes.UserId==100&& loginMes.UserPwd=="123456"{
	// 	fmt.Println("用户合法")
	// 	loginResMes.Code=200//合法
	// }else{
	// 	//fmt.Println("用户不合法")
	// 	loginResMes.Code =500 //不合法
	// 	fmt.Println("用户不存在，请注册")
	// 	loginResMes.Error = "用户不存在请注册"

	// }

	//loginResMes 序列化
	data,err:=json.Marshal(loginResMes)
	if err !=nil {
		fmt.Println("序列化失败",err)
		return
	}
	//将data 赋值给
	//将data 赋值给
	resMes.Data = string(data)
	//对他序列化
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("返回的数据序列化失败",err)
	}
	//发送data
	//因为使用分层要（mvc）。先创建实例
	tf :=&utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
func (this *UserProcess)ServiceProcessRegister(mes *message.Message)(err error){
	var registerMes message.RegisterMes
	err=json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=",err)
		return 
	}
	//先声明一个返回的结构体
	var resMes message.Message
	resMes.Type= message.RegisterMesType
	//声明一个LoginResM
	var RegisterMesType  message.RegisterResMes
	err =model.MyUserDao.Register(&registerMes.User)
	if err !=nil {
		if err == model.ERROR_USER_EXISTS{
			RegisterMesType.Code=505
			RegisterMesType.Error= model.ERROR_USER_EXISTS.Error()
		}else{
			RegisterMesType.Code=505
			RegisterMesType.Error= "未知错误"
		}
	}else{
		RegisterMesType.Code=200
	}
		data,err:=json.Marshal(RegisterMesType)
	if err !=nil {
		fmt.Println("序列化失败",err)
		return
	}
		resMes.Data = string(data)
	//对他序列化
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("返回的数据序列化失败",err)
	}
	//发送data
	//因为使用分层要（mvc）。先创建实例
	tf :=&utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}