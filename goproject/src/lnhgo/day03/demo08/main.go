package main
import (
	"fmt"
)
//定义
type Duck interface{
	//方法
	Gage()
	Walk()
	Swimming()
}
type pskDuck struct{
}
//接口实现
func (pd *pskDuck)Gage(){
	fmt.Println("gaga")
}
func (pd *pskDuck)Walk(){
	fmt.Println("Walk")
}
func (pd *pskDuck)Swimming(){
	fmt.Println("momg ")
}
func main(){
	//go 语言的接口，鸭子类型；到处都是鸭子类型duck typing
	/*
	当看到一只鸟走起来像鸭子，游泳起来像鸭子、叫起来像鸭子、那么这只鸟就是鸭子
	动词，方法，看是否具备方法；鸭子的行为是强调的外部行为
	*/
	var a Duck = &pskDuck{}
	a.Gage()
	
}