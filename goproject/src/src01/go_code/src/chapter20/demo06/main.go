package main

import (
	"fmt"

	//"github.com/redis/go-redis/v9/helper"
)
type CatNode struct{
	no int
	name string
	next *CatNode
	//pro *CatNode
}
func InsertCatNode(head *CatNode,newCatNode *CatNode){
//判断是不是diyizhimao
	if head.next==nil{
		head.no=newCatNode.no
		head.name=newCatNode.name
		head.next=head //构成一个环形
		//fmt.Println(newCatNode.name,"已加入环形链表")
		return
	}
	// 定义一个临时变量。帮忙找到环形的最后一个姐姐
	temp :=head
	for{
		if temp.next== head{
			break
		}
		temp=temp.next
	}
	//加入到环形列表
	//newCatNode.next=head
	newCatNode.next= head
	temp.next=newCatNode



}
func ListCircleLink(head *CatNode){
	temp := head
	if temp.next==nil{
		fmt.Println("空空如也")
		return
	}
	for{
		fmt.Printf("猫的信息为=[id=%d name=%s]\n",temp.no,temp.name)
		if temp.next==head{
			break
		}
		temp=temp.next
	}
}
//删除一致猫
//思路 先让temp 指向head；2让helper指向环形链表最后；3让temp 和要删除的di 进行比较；如果相同，则通过helper完成删除；要考虑一个情况
func DelCatNode(head *CatNode,id int) *CatNode{
	temp := head
	helper:= head 
	if temp.next==nil {
		fmt.Println("空链表无法删除")
		return head
	}
	//就一个节点
	if temp.next== head{
		temp.next=nil
		return head
	}
	//

	for {
		if helper.next==head{
			break
		}
		helper=head.next
	}
	flag:=true
	for{
		if temp.next==head{
			break
		}
		if temp.no==id{
			if temp ==head{//删除头节点

			}
			helper.next= temp.next
			//break
			flag= false
			break
		}
		temp = temp.next//比较
		helper = helper.next//价值找到删除节点
	}
	if flag {
		if temp.no== id{
			helper.next = temp.next
			fmt.Printf("猫咪=%d\n",id)
		}else{
			fmt.Println("对不起没有这只猫咪",id)
		}
	}
	return head


}
func  main(){
	head:=&CatNode{}
	cat1:=&CatNode{
		no: 1,
		name: "tom1",
	}
	// cat2:=&CatNode{
	// 	no: 2,
	// 	name: "tom2",
	// }
	// cat3:=&CatNode{
	// 	no: 3,
	// 	name: "tom3",
	// }
	// cat4:=&CatNode{
	// 	no: 4,
	// 	name: "tom4",
	// }
	InsertCatNode(head,cat1)
	//InsertCatNode(head,cat2)
	//InsertCatNode(head,cat3)
	//InsertCatNode(head,cat4)
	ListCircleLink(head)
	DelCatNode(head,30)
}
