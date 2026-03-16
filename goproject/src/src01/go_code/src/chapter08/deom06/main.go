package main

import (
	"fmt"
	"sort"
)
func main(){
	map1 :=make(map[int]int)
	map1[0]=0
	map1[1]=1
	map1[67]=9
	map1[5]=98
	fmt.Println(map1)
	//如何安装map的key顺序进行排序输出
	//1. 先将map的key 放到切片中
	//2.对切片进行排序
	//3遍历切片。按照k
	keys := []int{}
	for k := range map1{
		keys=append(keys,k)
	}
	fmt.Println(keys)
	fmt.Println("排序前打印map")
	for i,v := range map1{
		fmt.Printf("排序前map[%v]=%v ",i,v)
	}
	fmt.Println()

// 排序
    sort.Ints(keys)
	fmt.Println(keys)

	//排序打印map
	for _,v := range keys{
		fmt.Printf("排序后map[%v]=%v ",v,map1[v])
	}

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter08\deom05\main.go
// map[0:0 1:1 5:98 67:9]
// [0 1 67 5]
// 排序前打印map
// 排序前map[67]=9 排序前map[5]=98 排序前map[0]=0 排序前map[1]=1
// [0 1 5 67]
// 排序后map[0]=0 排序后map[1]=1 排序后map[5]=98 排序后map[67]=9
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter08\deom06\main.go
// map[0:0 1:1 5:98 67:9]
// [5 0 1 67]
// 排序前打印map
// 排序前map[0]=0 排序前map[1]=1 排序前map[67]=9 排序前map[5]=98
// [0 1 5 67]
// 排序后map[0]=0 排序后map[1]=1 排序后map[5]=98 排序后map[67]=9
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter08\deom06\main.go
// map[0:0 1:1 5:98 67:9]
// [0 1 67 5]
// 排序前打印map
// 排序前map[1]=1 排序前map[67]=9 排序前map[5]=98 排序前map[0]=0
// [0 1 5 67]
// 排序后map[0]=0 排序后map[1]=1 排序后map[5]=98 排序后map[67]=9