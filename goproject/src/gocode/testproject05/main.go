package main
import "fmt"
func main(){
    //+加号
	//1.正数2.相加操作3.字符串拼接
	var n1 int = +10
	fmt.Println(n1)
	var n2 int = 10 + 2
	fmt.Println(n2)
	var s1 string = "abc" + "def"
	fmt.Println(s1) 
	// / 除号
	fmt.Println(10/3)
	fmt.Println(10.0/3)
	//取模%
	fmt.Println(10%3)
	//++
	var a int = 10
	a++ 
	fmt.Println(a)
	//--自减
	var b int = 10
	b-- 
	fmt.Println(b)
	//++自增加1操作，--自减减1操作
	// go语言里++ -- 操作非常简单，只能单独使用，不能参与到运算
	//go语音里面,++ -- 只能在变量后面，不能在变量前面 --a ++a 不允许的
	var n3 int = 10
    fmt.Println(n3)
    var n4 int = (10+4) % 3 + 3 -7
	fmt.Println(n4)
	var n5 int = 10
	n5 += 20 //等价n5 = n5+20
	fmt.Println(n5)
	//练习
	var c int = 8
	var d int = 9
	fmt.Printf("c=%v,b=%v\n",a,b)
	var t int
	t = c
	c = d
	b = t
    fmt.Printf("c=%v,d=%v\n",a,b)
} 