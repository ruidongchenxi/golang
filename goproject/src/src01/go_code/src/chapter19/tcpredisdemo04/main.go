package main
import (
	"github.com/redis/go-redis/v9"
	"fmt"
	"context"
	//"time"

)
func main(){
	// var  name string
	// var  age int
	// var skill string
	// fmt.Println("输入名字")
	// fmt.Scanln(&name)
	// fmt.Println("输入年龄")
	// fmt.Scanln(&age)
	// fmt.Println("职业")
	// fmt.Scanln(&skill)
	ctx := context.Background()
	//通过go 向redis 写数据 
	conn:=redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379", //redis 地址
	    Password: "", //密码
	    DB: 0, 
		PoolSize: 6, //最大连接
		MinIdleConns: 2, //空闲最少连接
		}) //连的库
	//测试连接
	_, err := conn.Ping(ctx).Result()
	if err != nil {
		panic(err)
		//return
	}
	fmt.Println("连接 Redis 成功 ✅")
	

 }
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter19\tcpredisdemo02\main.go
// 输入名字
// chenxi
// 输入年龄
// 15
// 职业
// java
// 连接 Redis 成功 ✅
// 这次新写入的字段数量 3
// var3[name]=chenxi
// var3[age]=15
// var3[skill]=java