package main

import (
	"fmt"
	"math/rand"
	"sort"
)
type Hero struct{
	Name string
	Age int 
}
type HeroSlice []Hero

func (hs HeroSlice)Len() int{
	return len(hs)
}
// 按年龄从到大
func (hs HeroSlice)Less(i,j int) bool{
	return hs[i].Age>hs[j].Age
}
func (hs HeroSlice)Swap(i,j int){
	hs[i],hs[j]=hs[j],hs[i]
}
func main(){
	var intSlice = []int{1,9,0,-1,7,90}
	// 排序
	//使用系统
	fmt.Println(intSlice)
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	var t HeroSlice
	//t = make([]Hero, 10)
	for i:=0;i<10;i++{
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d",rand.Intn(100)),
			Age: rand.Intn(50+20),
		}
		t=append(t, hero)
	}
	fmt.Println(t)
	// for _,v := range t{
	// 	fmt.Println(v)
	// }
	sort.Sort(t)
	fmt.Println(t)
}
// 执行结过
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter11\demo03\main.go
// [1 9 0 -1 7 90]
// [-1 0 1 7 9 90]
// [{英雄~73 65} {英雄~43 16} {英雄~6 57} {英雄~19 23} {英雄~96 48} {英雄~27 67} {英雄~5 47} {英雄~20 55} {英雄~12 46} {英雄~55 28}]
// [{英雄~43 16} {英雄~19 23} {英雄~55 28} {英雄~12 46} {英雄~5 47} {英雄~96 48} {英雄~20 55} {英雄~6 57} {英雄~73 65} {英雄~27 67}]