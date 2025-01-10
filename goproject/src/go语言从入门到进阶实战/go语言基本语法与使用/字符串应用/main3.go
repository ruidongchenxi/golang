package main
import (
	"fmt"
)
//go语言的字符串无法直接修改每一个字符元素，只能通过重新构造的字符串并赋值给原来的字符串变量实现
func main()  {
	angel := "Heros never die"
	//将字符串转为字符串数组
	angelBytes := []byte(angel)
	//利用循环，将never单词替换为空
	for i := 5; i <=10; i++{
		angelBytes[i] = ' '
	}
	//打印替换后的结果
	fmt.Println(string(angelBytes))
	
}
//go 语言中的字符串和其他高级语言一样，默认是不可变。字符串不可变有很多好处，如天生线程安全,大家使用的都是只读对象，无需枷锁，再者，方便内存共享，而不必使用写时复制，字符串hash值也只需要制作一份
//所以说，代码中实际修改的是[]byte,[]byte在go语言里可变的，本身就是一个切片
//在完成了对[]byte操作后，在第9行，使用string()将[]byte转为字符串，重新创造了一个新的字符串
/*总结
    go 语言的字符串是不可变的
	修改字符串时，可以将字符串转换为[]byte进行修改
	[]byte和string可以通过强制类型转换互转
*/
