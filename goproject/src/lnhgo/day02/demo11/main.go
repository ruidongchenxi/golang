package main

import (
	"errors"
	"fmt"
)
func A() (int,error){
	//panic("异常退出")//会导致程序退出，不要随便用
	defer func ()  {
		if r := recover();r!=nil{
			fmt.Println("recover if A.",r)
		}
	}()
	var name map[string]string
	name["go"]="go"
	fmt.Println("this is a func")
	fmt.Println("nihsao")
	return 0, errors.New("失败")
}
func main(){
	_,err:=A()
	if err !=nil{
		fmt.Println(err)
	}

}