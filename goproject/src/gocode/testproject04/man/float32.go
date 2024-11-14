package main
import (
	"fmt"
)
func main(){
   // 定义字符类型的数据
   var c1 byte = 'a'
   fmt.Println(c1)
   var c2 byte = '('
   fmt.Println(c2)
   //字符类型：本质上是一个整数也可以参与运算
   //字母数字标点等字符也是ASCII 字符
   //var c3 byte = '中'
   //fmt.Println(c3)
   /*
   
# command-line-arguments
.\float32.go:13:8: c3 declared and not used
.\float32.go:13:18: cannot use '中' (untyped rune constant 20013) as byte value in variable declaration (overflows)
   */
   //汉字字符。底层对应的是Unicode
   //对应的码值为20013.byte类型溢出，能存储的范围可以用int
   //总结： golang 的字符对应的使用的是UTF-8编码（Unicode是对应的字符集）
   var c4 = '中'
   //显示对应字符必须采用格式化输出
   fmt.Printf("c4 对应的具体字符 %c",c4)
}