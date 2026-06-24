package main

import (
	"fmt"
	//"go/types"
)
func add(a,b int) int{
	return  a+b
}
func adds(a,b interface{}) int{
	ai,_:=a.(int)//断言
	bi,_:=b.(int)//断言
	return  ai+bi
}
func addsq(a,b interface{}) int{
	ai,ok:=a.(int)
	if !ok{
		panic("not an int type")
	}
	bi,ok:=b.(int)
	if !ok{
		panic("not an int type")
	}
	return ai+bi
}
func addsw(a,b interface{}) interface{}{
	switch a.(type){
	case int:
		ai,_:=a.(int)
		bi,_:=b.(int)
		return  ai+bi
	case float32:
		ai,_:=a.(float32)
		bi,_:=b.(float32)
		return  ai+bi
	case float64:
		ai,_:=a.(float64)
		bi,_:=b.(float64)
		return ai+bi
	case string:
		ai,_:=a.(string)
		bi,_:=b.(string)
		return ai+bi
	default:
		panic("ytr")
	}
}
func main(){
	a:=1.3
	b:=2.5
	fmt.Println(addsq(a,b))
}