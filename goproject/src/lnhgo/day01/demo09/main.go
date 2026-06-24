package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main(){
	name := "imooc体系"
	bytes := []rune(name)
	fmt.Println(len(bytes))
	//格式化输出
	username:="bobby"
	age := 18
	address :="北京"
	mobile := "1857656899"
	fmt.Printf("用户名:%s,年龄:%d,地址:%s,电话:%s\n",username,age,address,mobile)//极难维护
	usrMsg:=fmt.Sprintf("用户名:%s,年龄:%d,地址:%s,电话:%s",username,age,address,mobile)//
	fmt.Println(usrMsg)
	//通过string的builder进行字符串拼接，高性能
	var builder strings.Builder
	builder.WriteString("用户名:")
	builder.WriteString(username)
	builder.WriteString(",年龄:")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString(",地址:")
	builder.WriteString(address)
	builder.WriteString(",电话:")
	builder.WriteString(mobile)
	fmt.Println(builder.String())
}