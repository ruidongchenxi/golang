package main
import (
	"github.com/redis/go-redis/v9"
	"fmt"
	"context"
	//"time"

)
func main(){
	var  name string
	var  age int
	var skill string
	fmt.Println("输入名字")
	fmt.Scanln(&name)
	fmt.Println("输入年龄")
	fmt.Scanln(&age)
	fmt.Println("职业")
	fmt.Scanln(&skill)
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
	// err = conn.Set(ctx,"name","chenxi",0).Err()
	// if err !=nil{
	// 	panic(err)
	// }
	// err= conn.HSet(ctx,"users",map[string]interface{}{
	// 	"name1": name,
	// 	"age": age,
	// 	"job": skill,
	// }).Err()

	/*
	pipe.MGet(...) 没有执行，所以拿不到数据
	二、Pipeline 的本质（你必须理解）

👉 Pipeline 就像一个“命令缓冲区”：

你写的命令 → 先放进去（不会立刻执行）
只有 Exec() → 才真正发给 Redis
👉 这个 .Result() 的前提是：

命令已经执行过（Exec）

但你没有 Exec，所以：

没有返回值 ❌
可能是 nil / 空 ❌
1️⃣ pipe.xxx()   （收集命令）
2️⃣ pipe.xxx()   （继续收集）
3️⃣ Exec()       （统一执行）
4️⃣ cmd.Val()    （取结果）
七、你刚刚这个错误，本质说明什么？

👉 你已经开始进入：

“异步/批量执行模型”

这是 Go + Redis 的高级用法，不是初学者内容了。
八、再给一个进阶建议（非常重要）

以后写 pipeline，养成这个习惯：

cmd := pipe.Get(ctx, "key")

👉 不要直接 .Result()
👉 一定要：

cmd.Val()

在 Exec 之后取值
最后一句话（点醒你）

👉 Pipeline 不是“立即执行”，而是：

先收集 → 再统一执行 → 再取结果
	*/
	pipe := conn.Pipeline()
	var3:=pipe.HSet(ctx,"user",map[string]interface{}{
		"name":name,
		"age":age,
		"skill":skill,
	})
	ge:=pipe.HGetAll(ctx,"user")
	_, err = pipe.Exec(ctx)
	if err !=nil{
		fmt.Println("你好")
	}
	fmt.Println("这次新写入的字段数量",var3.Val())
	for i,k := range ge.Val(){
		fmt.Printf("var3[%v]=%v\n",i,k)
	}
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