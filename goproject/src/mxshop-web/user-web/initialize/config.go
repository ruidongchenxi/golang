package initialize

import (
	"fmt"

	//"mxshop-web/user-web/config"
	"mxshop-web/user-web/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
		debug:=GetEnvInfo("MXSHOP_DEBUG")//如果是刚设置的环境变量要重启开发工具像vscode、goland 等才会生效，因为启动的时候才会读环境变量，而你启动后设置的环境变量需要重启更新环境变量
	configFilePrefix:="config"
	cofigFileName:=fmt.Sprintf("user-web/%s-pro.yaml",configFilePrefix)
	if debug{
		cofigFileName = fmt.Sprintf("user-web/%s-debug.yaml",configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(cofigFileName)//配置相对路径
	if err:= v.ReadInConfig(); err!=nil{
		panic(err)
	}
	// server := config.ServerConfig{}
	//使用全局变量
	if err := v.Unmarshal(global.ServerConfig);err!=nil{
		panic(err)
	}
	//fmt.Println(global.ServerConfig)
	zap.S().Infof("配置信息:&v",global.ServerConfig)
	fmt.Printf("%v",v.Get("environment"))
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {

		fmt.Println("config file channed :",e.Name)
		zap.S().Infof("配置文件产生变化:%v",e.Name)
		_=v.ReadInConfig()
		_=v.Unmarshal(global.ServerConfig)
		// fmt.Println(global.ServerConfig)
		
		zap.S().Infof("配置信息:&v",global.ServerConfig)
	})
	///time.Sleep(time.Second*300)
}