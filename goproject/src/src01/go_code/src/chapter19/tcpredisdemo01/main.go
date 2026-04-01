package main
import (
	"github.com/redis/go-redis/v9"
	"fmt"
	"context"

)
func main(){
	ctx := context.Background()
	//通过go 向redis 写数据 
	conn:=redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379", //redis 地址
	    Password: "", //密码
	    DB: 0, }) //连的库
	//测试连接
	_, err := conn.Ping(ctx).Result()
	if err != nil {
		panic(err)
		//return
	}
	fmt.Println("连接 Redis 成功 ✅")
	//写数据
	err = conn.Set(ctx,"name","chenxi",0).Err()
	if err !=nil{
		panic(err)
	}
	err= conn.HSet(ctx,"users",map[string]interface{}{
		"name1": "小鹿",
		"age": 18,
		"job": "golang",
	}).Err()
	// 4. 读数据
	val, err := conn.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	//读数据
	val1,err := conn.HMGet(ctx, "users","name1","age","job").Result()
	fmt.Println(val1)
	
 }
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter19\tcpredisdemo01\main.go
// 连接 Redis 成功 ✅
// chenxi
// [小鹿 18 golang]