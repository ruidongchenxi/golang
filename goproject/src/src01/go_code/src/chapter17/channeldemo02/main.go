package main

import "fmt"
type cat struct{
	Name string
	Age int
}
func main(){
	// var mapChan chan map[string]string
	// mapChan = make(chan map[string]string,10)
	// m1:=make(map[string]string,10)
	// m1["city1"]= "北京"
	// m1["city2"]="天津"
	// m2:=make(map[string]string)
	// m2["hero1"]="宋江"
	// m2["hero2"]="武松"
	// mapChan <-m1
	// mapChan <-m2
	// fmt.Printf("长度=%v，容量=%v\n",len(mapChan),cap(mapChan))
	// //<-mapChan
	// fmt.Printf("m1=%v\n",<-mapChan)//取出m1的map直接格式化输出
	// fmt.Printf("长度=%v，容量=%v\n",len(mapChan),cap(mapChan))
	// var catChan chan cat
	// catChan = make(chan cat,10)
	// cat1:=cat{"小鹿",56}
	// cat2:=cat{"小狗",16}
	// catChan <-cat1 //写入
	// catChan<- cat2 //写入
	// fmt.Printf("长度=%v，容量=%v\n",len(catChan),cap(catChan))
	// fmt.Printf("cat1=%v\n",<-catChan)
	// fmt.Printf("长度=%v，容量=%v\n",len(catChan),cap(catChan))
	// var catChan chan *cat
	// catChan = make(chan *cat,10)
	cat1 := cat{"lu",16}
	// cat2 := cat{"t1",17}
	// catChan<-&cat1
	// catChan<-&cat2
	// cat11:=<-catChan
	// fmt.Printf("长度=%v，容量=%v\n",len(catChan),cap(catChan))
	// fmt.Println(*cat11)
	// fmt.Println(*(<-catChan))
	// fmt.Printf("长度=%v，容量=%v\n",len(catChan),cap(catChan))
	var allChan chan interface{}
	allChan= make(chan interface{},10)
	allChan<-cat1
	allChan<-&cat1
	allChan<-"晨曦"
	allChan<-78.6
	allChan<-89
	fmt.Printf("长度=%v，容量=%v\n",len(allChan),cap(allChan))
	a := (<-allChan).(cat)//这里类型断言
	fmt.Println(a.Name)//已经类型断言转换成cat结构所有可以直接调Name字段
	fmt.Println((<-allChan).(*cat).Name) //这里取Name值需要类型断言 
	fmt.Println(<-allChan)
	fmt.Printf("长度=%v，容量=%v\n",len(allChan),cap(allChan))


}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo02\main.go
// 长度=5，容量=10
// lu
// lu
// 晨曦
// 长度=2，容量=10