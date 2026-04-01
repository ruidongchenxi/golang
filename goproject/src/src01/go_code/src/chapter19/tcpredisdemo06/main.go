package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)
var (
	rdb *redis.Client
	ctx context.Context
)
func init(){
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		PoolSize:     10, // 控制并发连接
		MinIdleConns: 3,
	})
}

func main() {

	// 先写入数据
	//rdb.Set(ctx, "name", "chenxi", 0)
	pipe := rdb.Pipeline()

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("product:%d", i)

		pipe.HSet(ctx, key, map[string]interface{}{
			"name":  fmt.Sprintf("商品%d", i),
			"price": i * 10,
		})
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
	var cursor uint64
	for  {
	keys,nextCursor,err:=rdb.Scan(ctx ,cursor,"product:*",100).Result()
	if err !=nil{
		fmt.Println("错误")
		return
	}
	for _,key:=range keys{
		fmt.Println("key:", key)
		// 获取数据
		data, _ := rdb.HGetAll(ctx, key).Result()
		fmt.Println(data)
	}
	cursor = nextCursor

	if cursor == 0 {
		break
	}
	}
	
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter19\tcpredisdemo06\main.go
// key: product:13
// map[name:商品13 price:130]
// ...
// key: product:122
// map[name:商品122 price:1220]
