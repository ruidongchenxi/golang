package main

import (
	"context"
	"fmt"
	"sync"
	//"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		PoolSize:     10, // 控制并发连接
		MinIdleConns: 3,
	})

	// 先写入数据
	rdb.Set(ctx, "name", "chenxi", 0)

	var wg sync.WaitGroup

	// 模拟100个并发请求
	for i:= 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			val, err := rdb.Get(ctx, "name").Result()
			if err != nil {
				fmt.Println("错误:", err)
				return
			}

			fmt.Printf("goroutine %d -> %s\n", i, val)
		}(i)
	}

	wg.Wait()
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter19\tcpredisdemo05\main.go
// goroutine 9 -> chenxi
// goroutine 8 -> chenxi
// goroutine 3 -> chenxi
// goroutine 2 -> chenxi
// goroutine 4 -> chenxi
// goroutine 1 -> chenxi
// goroutine 6 -> chenxi
// goroutine 0 -> chenxi
// goroutine 7 -> chenxi
// goroutine 5 -> chenxi