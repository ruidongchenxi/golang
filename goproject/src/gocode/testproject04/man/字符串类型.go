package main
import "fmt"
func main(){
	var s1 string = "你好辛苦了"
	fmt.Println(s1)
	//2.字符串不可变的。指字符串一旦。其中字符不可以变
	var s2 string = "asde"
	//s2[0] = b
	fmt.Println(s2)
	//3.字符串表示形式
	//(1)如果字符串中没有特殊字符。字符串的表示形式用双引号没有问题
	var s3 string = "qwerfdsa"
	fmt.Println(s3)
	//(2)如果字符串有特殊字符就用反引号``
	var s4 string = `
	var flag01 bool = false
	var s1 string = "你好辛苦了"
	fmt.Println(s1)
	//2.字符串不可变的。指字符串一旦。其中字符不可以变
	var s2 string = "asde"
	s2[0] = b
	fmt.Println(s2)
	//3.字符串表示形式
	//(1)如果字符串中没有特殊字符。字符串的表示形式用双引号没有问题
	var s3 string = "qwerfdsa"
	fmt.Println(s3)
	`
    fmt.Println(s4)
}