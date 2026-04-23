package main

import (
	//"crypto/cipher"
	"fmt"
	//"github.com/redis/go-redis/v9/helper"
)
type Boy struct{
	No int
	Next *Boy
}
//编写函数，构成环形链表
//num 表示小孩个数
//Boy:返回该环形的链表的第一个小孩指针
func AddBoy(num int) *Boy{
	CurBoy := &Boy{}
	first := &Boy{}//空
	if num<1{
		fmt.Println("num的值")
		return first
	}
	//循环构建环形链表
	for i:=1;i<=num;i++{
		boy := &Boy{
			No: i,
		}
		//构成循环列表需要赋值指针

		//因为第一个小孩比较特殊
		if i==1{
			//boy.Next=boy
			first = boy
			CurBoy = boy
			CurBoy.Next = first//形成一个循环
		}else{
			CurBoy.Next=boy
			//CurBoy.Next.No=CurBoy.No+1
			CurBoy= boy
			CurBoy.Next=first
		}
	}
	return  first
}
//分析思路
// 1.编写一个函数


func ShowBoy(first *Boy){
	if first.Next==nil {
		fmt.Println("是一个空链表")
		return
	}
	//创建指针帮助遍历
	CurBoy := first
	for{
		// if CurBoy.Next==first{
		// 	fmt.Printf("小孩编号=%d\n",CurBoy.No)
		// 	break
		// }else{
		// 	fmt.Printf("小孩编号=%d\n",CurBoy.No)
		// 	CurBoy=CurBoy.Next
		// }
		fmt.Printf("小孩编号=%d\n",CurBoy.No)
		if CurBoy.Next== first{
			break
		}
		CurBoy= CurBoy.Next
	}	
	

}
//编写一个函数，PlayGame(first *Boy,startNo int,countNum int)
//最后使用环形链表留下最后一个人

func PlayGame(first *Boy,startNo int,countNum int){
	//需要定辅助节点。 帮助帮助我们删除删除小孩的一个，这是刚才说还有一个情况也是这样的，如果上来过这个空，如果是空的那个空的表就可以了，有问题空表呢，单独一下空的链表，链表，我们单独的出一下单独处理，这个很简单，那就意思，如果我们进来投机里的，那他等于is near说明它是空的列表就没法没法玩，所以放一个控制好。
	//如果这个是一个空列表。 我们就要单独处理
	if first.Next==nil{
		fmt.Println("空")
		return
	}
	//留一个判断。 参数是否大于小孩总数
	tail := first
	// Tail指向环形链表的最后一个。
	//因为tail在删除小孩子使需要用到。
	for{
		if tail.Next==first{
			break
		}
		tail=tail.Next
	}
	for i:=1;i<= startNo-1;i++{
		first = first.Next
		tail= tail.Next
	}
	for {
		// 开始数countNum-1
		for i:=1; i<=countNum-1;i++{
			first = first.Next
			tail= tail.Next
		}
		//删除first 执行的小孩
		fmt.Printf("小孩编号为%d\n",first.No)
		first = first.Next
		tail.Next = first
		if first== tail{
			break
		}
	}
	fmt.Printf("小孩编号为%d 出圈->\n",first.No)
 }

func main(){
	first := AddBoy(5)
	ShowBoy(first)
	PlayGame(first,2,3)
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter20\demo07\main.go
// 小孩编号=1
// 小孩编号=2
// 小孩编号=3
// 小孩编号=4
// 小孩编号=5
// 小孩编号为4
// 小孩编号为2
// 小孩编号为1
// 小孩编号为3
// 小孩编号为5 出圈->