package main

import (
	//"fmt"
	"fmt"
	"testing" //引入框架包
)

//编写测试用例测试
func TestAddUpper(t *testing.T){
	res :=addUpper(10) 
	if res!=55{
		// fmt.Printf("addupper错误 返回值=%v，期望值=%v\n",res,55)
		t.Fatalf("addupper错误 返回值=%v，期望值=%v\n",res,55)
	}
	t.Logf("addupper执行正确")

}
func TestHello(t *testing.T) {
	fmt.Println("testHell调用....")
}
func TestKaa(t *testing.T) {
	res:=getSub(50,23)
	fmt.Println("差",res)
	
}