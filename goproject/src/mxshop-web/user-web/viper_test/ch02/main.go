package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)
type MysqlConfig struct{
	Host string `mapstructure:"host"`
	Prot  int `mapstructure:"prot"`
}
type ServerConfig struct{
	Environment string `mapstructure:"environment"`
	ServerName string `mapstructure:"name"`//tag 表示把mapstructure的name对应的值填充到ServerName
	Prot int `mapstructure:"prot"`//
	MysqlInfo MysqlConfig `mapstructure:"mysql"`

}
func GetEnvInfo(env string) bool{
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {
	debug:=GetEnvInfo("MXSHOP_DEBUG")//如果是刚设置的环境变量要重启开发工具像vscode、goland 等才会生效，因为启动的时候才会读环境变量，而你启动后设置的环境变量需要重启更新环境变量
	configFilePrefix:="config"
	cofigFileName:=fmt.Sprintf("user-web/viper_test/ch02/%s-pro.yaml",configFilePrefix)
	if debug{
		cofigFileName = fmt.Sprintf("user-web/viper_test/ch02/%s-debug.yaml",configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(cofigFileName)//配置相对路径
	if err:= v.ReadInConfig(); err!=nil{
		panic(err)
	}
	server := ServerConfig{}
	if err := v.Unmarshal(&server);err!=nil{
		panic(err)
	}
	fmt.Println(server)
	fmt.Printf("%v",v.Get("environment"))
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file channed :",e.Name)
		_=v.ReadInConfig()
		_=v.Unmarshal(&server)
		fmt.Println(server)
	})
	time.Sleep(time.Second*300)
}
