package main
import (
	_ "fmt"
)
func main(){
	// 创建一个int到int的映射
	m := make(map[int]int)
	//开启一段并发代码
	go func() {
		//不停地对map进行
		for {
			m[1] = 1
		}
	}()
	//开启一段并发代码
	go func(){
		// 不停地对map进行读取
		for {
			_ = m[1]
		}
	}()
	//无限循环，让并发程序在后台执行
	for {

	}
}
/*执行报错
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\映射\main5.go
fatal error: concurrent map read and map write

goroutine 7 [running]:
main.main.func2()
        D:/golang/goproject/src/go语言从入门到进阶实战/容器/映射/main5.go:19 +0x28
created by main.main in goroutine 1
        D:/golang/goproject/src/go语言从入门到进阶实战/容器/映射/main5.go:16 +0x96

goroutine 1 [runnable]:
main.main()
        D:/golang/goproject/src/go语言从入门到进阶实战/容器/映射/main5.go:23 +0x96

goroutine 6 [runnable]:
main.main.func1()
        D:/golang/goproject/src/go语言从入门到进阶实战/容器/映射/main5.go:12 +0x28
created by main.main in goroutine 1
        D:/golang/goproject/src/go语言从入门到进阶实战/容器/映射/main5.go:9 +0x58
exit status 2
运行时输出提示： 并发的map读写，也就是说使用了两个并发函数不断地对map进行读和写而发生了竟态问题。map内部会对这种并发操作进行检查并提前发现。
需要并发读写时，一般的做法是加锁，但这样性能并不高。go语言在1.9版本中提供了一种效率较高的并发安全的sync.Map。sync.Map和map不同，不是以语言原生形态提供，而是在sync包下的特殊结构

*/