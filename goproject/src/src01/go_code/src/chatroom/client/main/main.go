package main

import (
	"fmt"
	"os"
	"src/chatroom/client/process"
)

//定义两个全局变量一个表示用户ID；一个表示用户密码
var (
	userId int
	userPwd string
	userName string
)
func main(){
	//接收用户选择
	var key int
	//判断是否还继续显示菜单
	//var loop =true
	for {
		fmt.Println("------------------------------欢迎登录多人聊天系统-----------------------------")
		fmt.Println("\t\t\t1.登录聊天室")
		fmt.Println("\t\t\t2.注册用户")
		fmt.Println("\t\t\t3.退出系统")
		fmt.Printf("请选择: ")
		fmt.Scanf("%d\n",&key)
		switch key{
			case 1:
				fmt.Println("登录聊天室")
				fmt.Println("请输入ID：")
				fmt.Scan(&userId)
				fmt.Println("输入用户密码")
				fmt.Scan(&userPwd)
				//
				up:=&process.UserProcess{}
				up.Login(userId,userPwd)

			case 2:
				fmt.Println("注册用户")
				fmt.Println("请输入id")
				fmt.Scan(&userId)
				fmt.Println("输入密码")
				fmt.Scan(&userPwd)
				fmt.Println("输入名字")
				fmt.Scan(&userName)
				//loop=false
				up:=&process.UserProcess{}
				up.Register(userId,userPwd,userName)
			case 3:
				fmt.Println("退出聊天室")
				//loop=false
				os.Exit(0)
			default:
				fmt.Println("输入有误重新输入")
		}
	}
}