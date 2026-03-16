package main
import (
	"fmt"

)
func main(){
	a := make(map[string]string)
	a["a1"]= "阿库"
	a["a2"]= "小米"
	a["a3"]= "小鹿"
	a["a4"]= "小鹏"
	//遍历k
	// fmt.Println("遍历k")
	// for k:= range a{
	// 	fmt.Println(k)
	// }
	// //遍历整个几个
	// fmt.Println("遍历整个集合")
	// for v,k := range a{
	// 	fmt.Println(v,k)
	// }
	fmt.Println("map a 的长度",len(a))
	// cities5 := make(map[string]map[string]string)
	// //cities5= make(map[string]map[string]string)
	// cities5["学生1"]= make(map[string]string)//这句话不能少
	// cities5["学生2"]= make(map[string]string)
	// cities5["学生3"]= make(map[string]string)
	// cities5["学生1"]["名字"]= "tom"
	// cities5["学生1"]["性别"]= "男"
	// cities5["学生1"]["地址"]= "胡同"
	// cities5["学生2"]["名字"]= "丽娜"
	// cities5["学生2"]["性别"]= "女"
	// cities5["学生2"]["地址"]= "天通苑"
	// cities5["学生3"]["名字"]= "杰克"
	// cities5["学生3"]["性别"]= "男"
	// cities5["学生3"]["地址"]= "青年汇"
	// //fmt.Println(cities5["学生3"])
	// for x := range cities5{
	// 	for i,v := range cities5[x]{
	// 		fmt.Printf("学生%v的%v是%v\n",x,i,v)
	// 	}
	// }

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter08\deom04\main.go
// map a 的长度 4