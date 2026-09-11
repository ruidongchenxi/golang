package main

import (
	"fmt"
	"mxshop-web/user-web/global"
	"mxshop-web/user-web/initialize"

	"go.uber.org/zap"
	//"google.golang.org/protobuf/proto"
)

func main() {
	// prot := 8021
	//初始化log
	initialize.InitLogger()
	//初始化配置文件
	initialize.InitConfig()
	//初始化路由
	Router :=initialize.Routers()
	


	// logger,_:=zap.NewDevelopment()
	// defer logger.Sync()
	// suger := logger.Sugar()
	/*
	1.S()可以获取一个全局的Sugar,可以设置一个全局的loggar
	2.日志是分级别的debug、info、 warn、 error 、fatal
	3.S函数和L函数很有用；给我们提供一个安全的全局访问路径
	*/
	zap.S().Debugf("启动，端口：%d",global.ServerConfig.Prot)//打印日志信息
	if err :=Router.Run(fmt.Sprintf(":%d",global.ServerConfig.Prot));err!=nil{
		zap.S().Panic("启动失败",err.Error())
		
	}
}
