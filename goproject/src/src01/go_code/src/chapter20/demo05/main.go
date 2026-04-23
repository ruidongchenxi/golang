package main

import (
	//"errors"
	"fmt"
	//"text/template"
	//"hash"
	//"os"
)

//node
type HeroNode struct{
	no int
	name string
	nickname string
	next *HeroNode
	pre *HeroNode
}
//删除节点
func DelHerNode(head *HeroNode,id int){
	temp:= head
	flag:= false
		for  {
		if temp.next ==nil{//到链表最后
			break
		}else if temp.next.no== id{
			//fmt.Println("错误不允许插入")
			flag = true
			break
		}
		temp= temp.next
	}

	if  flag {
		temp.next= temp.next.next
		if temp.next!=nil {
			temp.next.pre=temp
		}
		
	}else {
		fmt.Println("删除ID不存在")
	}

}
//给链表插入节点
// 给双向链表添加一个节点
func InsertHeroNode(head *HeroNode,h1 *HeroNode){
	//找到链表最后一个节点
	//创建一个辅助节点
	temp := head
	for {
		if temp.next==nil{// 表示找到最后
			break
		}
		temp = temp.next//不断指向下一个节点
	}
	temp.next= h1
	h1.pre= temp
}
func InsertHeroNode2(head *HeroNode,h1 *HeroNode){
	//找到链表适当节点
	//创建一个辅助节点
	temp := head
	flag := true
	for  {
		// if temp.next==nil{// 表示找到最后
		// 	break
		// }
		// temp = temp.next//不断指向下一个节点
		if temp.next ==nil{//到链表最后
			break
		}else if temp.next.no> h1.no{
			//说明h1 应该插入temp后面
			break
			//temp.next=h1
		}else if temp.next.no== h1.no{
			//fmt.Println("错误不允许插入")
			flag = false
			break
		}
		temp= temp.next
	}
	if  !flag {
		fmt.Println("对不起已经存在英雄")
		return
	}else{
		h1.next= temp.next
		//
		h1.pre=temp
		if temp.next!=nil {
		temp.next.pre=h1
		
		}
		temp.next= h1
	}
	temp.next= h1
	h1.pre= temp
}
//显示链表所有信息
func ListLink(head *HeroNode){
	//
	temp := head
	//判断链表是不是空链表
	if temp.next== nil{
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[%d ,%s,%s]===>\n",temp.next.no,temp.next.name,temp.next.nickname)
		temp= temp.next
		if temp.next== nil{
			fmt.Println("链表显示完成")
			break
		}	
	}
}
func ListLink2(head *HeroNode){
	//
	temp := head
	//判断链表是不是空链表
	if temp.next== nil{
		fmt.Println("空链表")
		return
	}
	for{
		if temp.next==nil{
			break
		}
		temp= temp.next
	}

	for {
		fmt.Printf("[%d ,%s,%s]===>\n",temp.no,temp.name,temp.nickname)
		temp= temp.pre

		if temp.pre== nil{
			fmt.Println("链表显示完成")
			break
		}	
	}
}

func main()	{
	head:= &HeroNode{
	}
	head1 := &HeroNode{
		no: 1,
		name: "宋江",
		nickname: "及时雨",

	}
	head2 := &HeroNode{
		no: 2,
		name: "卢俊义",
		nickname: "玉麒麟",

	}
	head3 := &HeroNode{
		no: 3,
		name: "林冲",
		nickname: "豹子头",

	}
	
	InsertHeroNode(head,head1)
	InsertHeroNode(head,head2)
	InsertHeroNode(head,head3)
	ListLink(head)

	ListLink2(head)
	fmt.Println("***")
}
	

