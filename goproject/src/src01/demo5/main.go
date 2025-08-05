package main
import "fmt"
func main(){
    var c1 byte = 'a'
	fmt.Println(c1)	//定义字符类型的数据
	var c2 byte = '6'
	fmt.Println(c2)
	var c3 byte = '('
	fmt.Println(c3)
	//字符类型本质上就是整数，也可以参与运算。输出时，会将对应的码值做一个输出
	//字符等是按照ASCII 码表
	var c4 int = '中' //汉字字符对应的码值，底层对应的是Unicode码值，对应的码值为20013，byte类型溢出，能存储范围：可以用int
	fmt.Println(c4)  //总结 golang的字符对应使用的是UTF-8编码
	var c5 byte = 'A'
	fmt.Printf("c5对应的具体字符：%c",c5)
}