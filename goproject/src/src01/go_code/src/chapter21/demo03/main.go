package main

import (
	//"embed"
	"fmt"
	"os"
	//"time"
	//"golang.org/x/text/cases"
)

//定义
type Emp struct{
	Id int 
	Name string
	Next *Emp
}
func (this *Emp)showMe(){
	fmt.Printf("链表%d,找到该雇员%d\n",this.Id % 7,this.Id)
}
//定义EmpLink，
type EmpLink struct{
	Head *Emp
}
func (this *EmpLink)Insert(emp *Emp){
	cur := this.Head
	var pre *Emp =nil
	//如果是空指针
	if cur == nil{
		this.Head=emp
		return
	}
	//如果不是一个空列表，给emp找到对应位置并插入
	//思路让cur 和emp 比较，然后pre保持在cru 前面
	for {
		if cur != nil{
			if cur.Id>=emp.Id{
				//找到位置
				break
			}
			pre =cur//
			cur= cur.Next
		}else{
			break
		}
	}
	// 退出时。将看一下是否将emp添加到链表的最后
	//if cur ==nil{
		pre.Next = emp
		emp.Next= cur
	// }else{
	// 	pre.Next = emp
	// 	emp.Next= cur
	// }
}
//显示当前链表信息
func (this *EmpLink)ShowLink(no int){
	if this.Head ==nil{
		fmt.Printf("链表%d 为空\n",no)
		return
	}
	cur := this.Head
	for{
		if cur !=nil{
			fmt.Printf("链表%d雇员id=%d名字=%s->",no,cur.Id,cur.Name)
			cur= cur.Next
		}else{
			break
		}
	}
	fmt.Println()
}
//根据ID查雇员
func (this *EmpLink) FindById(id int) *Emp{
	cur :=this.Head
	for{
		if cur !=nil&& cur.Id==id{
			return cur
		}else if cur == nil{
			break
		}
		cur = cur.Next
	}
	return nil
}
type HashTable struct{
	Head *Emp
	LinkArr [7]EmpLink

}
//给哈希TAble 编写Insert 雇员的方法。
func(this *HashTable)Insert(emp *Emp){
	linkNo:=this.HashFun(emp.Id)
	//确定将雇员添加哪个列表
	this.LinkArr[linkNo].Insert(emp)

} 
//显示所有雇员
func (this *HashTable)ShowAll(){
	for i:=0;i<len(this.LinkArr);i++{
		this.LinkArr[i].ShowLink(i)
	}
}
//编写用于散列函数
func (this *HashTable)HashFun(id int) int{
	return  id %7

}
//编写一个查找方法
func (this *HashTable)FindById(id int)*Emp{
	linkNo:=this.HashFun(id)
	return  this.LinkArr[linkNo].FindById(id)
}

func main(){
	var hashtable HashTable
	key := ""
	id := 0
	name := ""
	for {
		fmt.Println("============选择菜单================")
		fmt.Println("input 添加雇员")
		fmt.Println("show 显示所有雇员")
		fmt.Println("find 查找雇员")
		fmt.Println("exit 退出")
		fmt.Printf("输入你的选择：")
		fmt.Scanln(&key)
		switch key{
		case "input":
			fmt.Printf("输入雇员ID")
			fmt.Scanln(&id)
			fmt.Printf("输入名字")
			fmt.Scanln(&name)
			emp := &Emp{
				Name: name,
				Id: id,
			}
			hashtable.Insert(emp)

		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Printf("输入要查找的id号：")
			fmt.Scanln(&id)
			emp:=hashtable.FindById(id)
			if emp  ==nil{
				fmt.Println("用户不存在")
			}else{
				emp.showMe()
			}

		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}

	}

}
//查找
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter21\demo03\main.go
// ============选择菜单================
// input 添加雇员
// show 显示所有雇员
// find 查找雇员
// exit 退出
// 输入你的选择：input
// 输入雇员ID34
// 输入名字chenxi
// ============选择菜单================
// input 添加雇员
// show 显示所有雇员
// find 查找雇员
// exit 退出
// 输入你的选择：find
// 输入要查找的id号：34
// 链表6,找到该雇员34
// ============选择菜单================
// input 添加雇员
// show 显示所有雇员
// find 查找雇员
// exit 退出
// 输入你的选择：exit