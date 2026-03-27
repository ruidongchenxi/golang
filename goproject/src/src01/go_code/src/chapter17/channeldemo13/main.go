package main
import (
	"fmt"
	//"time"
)

func main(){
// 使用select 可以解决从管道取数据的阻塞问题
// 定义一个管道10个数据
intchan := make(chan int,10)
for i :=0;i <10; i ++{
	intchan <- i
}
strchan:=make(chan string,5)
for i := 0;i<5;i++{
	strchan <- "helloo"+fmt.Sprintf("%v",i)
}
//传统的方法遍历管道时不关闭会阻塞，导致死锁

//问题在实际开发中不好确定什么时候开关管道；可以使用selectl 方式也可以解决
t1:
for {
	select {
		case v:=<-intchan://注意：这里，如果intchan一直没有关闭，不会一直阻塞而导致死锁；会自动到下一个case匹配
			fmt.Println("从intchan读取了数据",v)
		case v:= <-strchan:
			fmt.Println("从strchan读取了数据",v)
		default:
			fmt.Println("都取不到数据，加入自己逻辑")
			break t1
		

	}
}

}

// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo13\main.go
// 从intchan读取了数据 0
// 从intchan读取了数据 1
// 从strchan读取了数据 helloo0
// 从intchan读取了数据 2
// 从intchan读取了数据 3
// 从strchan读取了数据 helloo1
// 从strchan读取了数据 helloo2
// 从strchan读取了数据 helloo3
// 从intchan读取了数据 4
// 从strchan读取了数据 helloo4
// 从intchan读取了数据 5
// 从intchan读取了数据 6
// 从intchan读取了数据 7
// 从intchan读取了数据 8
// 从intchan读取了数据 9
// 都取不到数据，加入自己逻辑