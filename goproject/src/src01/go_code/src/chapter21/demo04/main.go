package main

import (
	//"embed"
	"fmt"
	//"os"
	//"time"
	//"golang.org/x/text/cases"
)
type Hero struct{
	No int
	Name string
	Left *Hero
	Right *Hero

}
//前序遍历
func PreOrder(node *Hero){
	if node  != nil{
		fmt.Printf("no=%d name=%s\n",node.No,node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}
//中旬
func InfixOrder(node *Hero){
	if node  != nil{
		
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n",node.No,node.Name)
		InfixOrder(node.Right)
	}
}
func PostOrder(node *Hero){
	if node  != nil{
		
		PostOrder(node.Left)
		
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n",node.No,node.Name)
	}
}
func main(){
	root:=&Hero{
		No:1,
		Name: "宋江",

	}
	left1:= &Hero{
		No: 2,
		Name: "吴用",

	}
	left10:= &Hero{
		No: 10,
		Name: "tom",

	}
	left12:= &Hero{
		No: 12,
		Name: "杰克",

	}

	right1:=&Hero{
		No: 3,
		Name: "卢俊义",
	}
	root.Left=left1
	root.Right=right1
	right2:=&Hero{
		No: 4,
		Name: "林冲",
	}
	left1.Left=left10
	left1.Right=left12
	right1.Right=right2
	//PreOrder(root)
	//InfixOrder(root)
	PostOrder(root)

}
//结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter21\demo04\main.go
// no=10 name=tom
// no=12 name=杰克
// no=2 name=吴用
// no=4 name=林冲
// no=3 name=卢俊义
// no=1 name=宋江