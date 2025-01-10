package main
import (
	"fmt"
	"sync"
)
func main(){
	var scene sync.Map
	//将键值对保存到sync.Map
	scene.Store("greece",97)
	scene.Store("london",100)
	scene.Store("egypt",200)
	//从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	//根据键删除对应的键值对
	scene.Delete("london")
	//遍历所有sync.Map中的键值对
	scene.Range(func(k,v interface{}) bool{
		fmt.Println("iterate:", k,v)
		return true
	})
}
/* 执行结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\映射\main6.go                                  
100 true                                                                                                                
iterate: greece 97                                                                                                      
iterate: egypt 200                                                                                                      
声明scene,类型为sync.Map。注意，sync.Map不能使用make创建
将以系列键值对保存到sync.Map中。sync.Map将键和值以interface{}类型进行保存
提供一个sync.Map的键给scene.Load()方法后将查询到键对应的值返回
sync.Map的Delete可以使用指定的键将对应的键值对删除
Range()方法可以遍历sync.Map,遍历需要提供一个匿名函数，参数为k、v,类型为interface{}，每次Range()在遍历一个元素时，都会调用这个匿名函数把结果返回
	sync.Map没有提供获取map数量的方法，替代方法是获取时遍历自行计算数量。sync.Map为了保证并发安全有一些性能损失，因此在非并发情况下，使用map相比使用sync.Map会有更好的性能

*/