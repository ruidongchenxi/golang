package main

import "fmt"

func main(){
	var c1 byte = 'a'
	var c2 byte = '0'
	// 当我们直接输出byte值，就是输出了对应的字符的码值
	fmt.Println("c1=",c1)
	fmt.Println("c2=",c2)
	//如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1 = %c c2 = %c\n", c1,c2)
	var c3 int = '北'
	fmt.Println("c3=",c3)
	fmt.Printf("c3=%c\n",c3)
	var c4 int = 230
	fmt.Printf("c4=%c\n",c4)
	var c5 int = '陈'
	fmt.Println(c5+10)

}
// 执行结果
// c1= 97  
// c2= 48
//c1 = a c2 = 0
// c3= 21271
// c3=北
// c4=æ
// 38482