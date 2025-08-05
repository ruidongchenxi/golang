package main
import "fmt"
func main(){
	//定义
	var s1 string = "你好我是亚马逊"
	fmt.Println(s1)
	//字符串是不可变的
	var s2 string = "abc"
	//s2[0] = 't'
	fmt.Println(s2)
/* # command-line-arguments
demo8\mian.go:9:2: cannot assign to s2[0] (neither addressable nor a map index expression) */
    //字符串的表示形式,如果字符串没有特殊字符，字符串表示形式用双引号
	var s3 string = "as"
	//如果字符串有特殊字符，字符串表示形式用反引号
	var s4 string = `
	package main //声明包，每个go文件必须有归属包
import "fmt" //导入工具包，使用
func main(){ // main 主函数
	fmt.Println("你好")// 打印输出内容
	var age = 10 +9
	fmt.Println(age) 
	fmt.Println("qqqqqqqqqq",
	"qqqqqqqqqqqq",
	"qqqqqqqqqqqq",
	"qqqqqqqqqqqq",
	"qqqqqqqqqqqqq")
}
`
	fmt.Println(s3)
	fmt.Println()
	fmt.Println(s4)

	//字符串拼接
	var s5 string = "abc" + "def"
	s5 += "hijk"
	fmt.Println(s5)
	fmt.Println()

	//当一个字符串拼接过长 注意+ 保留在上一行最后
	var s6 string = "cx" + "xx" +
	"ff" + "qq" +
	"cd"
	fmt.Println(s6)

}