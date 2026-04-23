package main

import (
	"errors"
	"fmt"
	"strconv"
	//"golang.org/x/text/number"
	//"testing/quick"
)

//最左边下标left
//右边下标right
type Stack struct{
	MaxTop int //表示我们栈最大可以存放个数
	Top int //表示栈顶
	arr [20]int
}
func (this *Stack)push(val int) (err error){
	// 
	if this.Top==this.MaxTop-1{
		fmt.Printf("栈满")
		return errors.New("栈满溢出")
	}
	this.Top++
	this.arr[this.Top]=val
	return

}
func (this *Stack)Pop() (val int,err error){
	if this.Top==-1{
		fmt.Println("空栈")
		//val = -2
		return -2, errors.New("栈空")
	}
	val = this.arr[this.Top]
	this.Top--
	return val,nil
}
//遍历栈时需要从栈顶开始遍历
func (this *Stack)list(){
	//判断栈是否为空
	if this.Top==-1{
		fmt.Println("是一个空栈")
	}
//	curTop = this.Top
	for i:=this.Top;i>=0;i--{
		fmt.Printf("arr[%d]=%d\n",i,this.arr[i])
	}
}
//判断一个字符是不是运算符
func (this *Stack) IsOper(val int) bool{
	if val ==42||val==43 || val==45|| val==47{
		return true
	}else{
		return false
	}
}
func (this *Stack) Cal(num1 int,num2 int ,oper int) int{
	res :=0
	switch oper{
	case 42:
		res= num2*num1
	case 43:
		res=num2+num1
	case 45:
		res= num2-num1
	case 47:
		res= num2/num1
	default:
		fmt.Println("运算符错误")

	}
	return res
}
//编写一个方法返回某个运算符优先级
func (this *Stack) Priority(oper int) int{
	res := 0
	if oper==42|| oper==47 {
		res= 1
	}else if oper== 43|| oper==45{
		res= 0
	}
	return res


}
func main(){
	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top: -1,
	}
	operStack := &Stack{
		MaxTop: 20,
		Top: -1,
	}
	exp := "30+30*6-4"
	index := 0
	num1:=0
	num2:=0
	oper:= 0
	result:=0
	keepNum := ""
	for {
		//增加一个逻辑;处理
		ch := exp[index:index+1]//字符串
		temp := int([]byte(ch)[0])//字符对应aSSIC码值
		if operStack.IsOper(temp) {//说明
			if operStack.Top == -1{
				operStack.push(temp)
			}else{
				if operStack.Priority(operStack.arr[operStack.Top])>=operStack.Priority(temp){
					num1,_=numStack.Pop()
					num2,_=numStack.Pop()
					oper,_=operStack.Pop()
					result= operStack.Cal(num1,num2,oper)
					//结果压入数栈
					numStack.push(result)
					//结果压入符号栈
					operStack.push(temp)

				}else{
					operStack.push(temp)
				}
				
			}
			
		}else{//数字
			//多位数
			//1.定义一个变量，keepNum string,做拼接
			//2.每次要向后
			//如果到已经到表达最后，直接keepMum
			keepNum += ch
			if index == len(exp)-1{
				val,_ := strconv.ParseInt(keepNum,10,64)
				numStack.push(int(val))

			}else {
				if operStack.IsOper(int([]byte(exp[index+1:index+2])[0])){
					val,_ := strconv.ParseInt(keepNum,10,64)
					numStack.push(int(val))
					//operStack.push(int(val))
					keepNum=""

				}
			}

			// val,_ := strconv.ParseInt(ch,10,64)
			// numStack.push(int(val))
		}
		//继续扫描;判断index 是否到最后
		if index+1 == len(exp) {
			break
		}
		index++

	}
	for{
		if operStack.Top==-1{
			break
		}
		num1,_=numStack.Pop()
		num2,_=numStack.Pop()
		oper,_=operStack.Pop()
		result= operStack.Cal(num1,num2,oper)
		//结果压入数栈
		numStack.push(result)
		//结果压入符号栈
		//operStack.push(temp)

	}
  // 如果算法正确没有问题，表达式也是正确，则结果就是numStack 最后
  res,_:=numStack.Pop()
  fmt.Printf("表达式%s=%v\n",exp,res)


}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter20\demo12\main.go
// 表达式30+30*6-4=206