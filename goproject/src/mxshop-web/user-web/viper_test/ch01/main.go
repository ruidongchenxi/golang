package main

import (
	"fmt"

	"github.com/spf13/viper"
)
type ServerConfig struct{
	ServerName string `mapstructure:"name"`//tag 表示把mapstructure的name对应的值填充到ServerName
	Prot int `mapstructure:"prot"`//

}

func main() {
	v := viper.New()
	v.SetConfigFile("user-web/viper_test/ch01/config.yaml")//配置相对路径
	if err:= v.ReadInConfig(); err!=nil{
		panic(err)
	}
	server := ServerConfig{}
	if err := v.Unmarshal(&server);err!=nil{
		panic(err)
	}
	fmt.Println(server)

}
