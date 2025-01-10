package main
import (
	"fmt"
	"container/list"
)
func main(){
	l := list.New()
	l.PushBack("fist")
	l.PushFront("67")
	fmt.Println(*l)
}
/*
创建一个列表实例
将fish字符串插入到列表的尾部，此时列表是空的，插入后只有一个元素
将数值67放入列表。此时。列表中已经存在fist元素，67这个元素将被放在fist的前面
执行结果

PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\列表\main0.go
&{{0xc000020150 0xc000020120 <nil> <nil>} 2}
*/