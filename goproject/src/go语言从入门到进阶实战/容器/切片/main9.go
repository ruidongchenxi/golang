package main
import (
	"fmt"
)
func main(){
	seq := []string{"a","b","c","d","e"}
	//指定删除位置
	index := 2
	//查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index],seq[index+1:])
	//将删除点前后的元素拼接起来
	seq = append(seq[:index],seq[index+1:]...)
	fmt.Println(seq)
}
/* 执行后结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\切片\main9.go
[a b] [d e]
[a b d e]
声明一个整型切片，包含从a到e的字符串
为了演示和讲解，使用index变量保存需要删除的元素下标
seq[:index] 表示的就是被删除元素的前半部分只为[a,b]，seq[index+1:]表示的是被删除元素的后半部分，值为[d e]
使用append()函数将两个切片连接起来
输出连接好的新切片
*/