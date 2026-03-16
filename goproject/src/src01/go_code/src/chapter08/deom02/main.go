package main
import (
	"fmt"
)
func main(){
	//方式1
	//声明
	var cities1 map[string]string
	//使用make 初始化
	cities1= make(map[string]string)
	cities1["a"]= "golang"
	//方式2声明并初始化
	var cities2 = make(map[string]string)
	cities2["q"]="你好"
	//方式3 声明初始化并赋值（make没有明确写，但底层给make了）
	var cities3 map[string]string= map[string]string{"node1":"晨曦"}
	cities3["node2"]= "上海"
	fmt.Println("方式1",cities1)
	fmt.Println("方式2",cities2)
	fmt.Println("方式3",cities3)
	cities4 := map[string]string{
		"悟空":"金箍棒",
		"八戒":"九齿钉耙",
		"沙和尚":"降魔杖",
		"哪吒":"火尖枪",
	}
	fmt.Println("方式4",cities4)
	cities5 := make(map[string]map[string]string)
	//cities5= make(map[string]map[string]string)
	cities5["学生1"]= make(map[string]string)//这句话不能少
	cities5["学生2"]= make(map[string]string)
	cities5["学生3"]= make(map[string]string)
	cities5["学生1"]["名字"]= "tom"
	cities5["学生1"]["性别"]= "男"
	cities5["学生1"]["地址"]= "胡同"
	cities5["学生2"]["名字"]= "丽娜"
	cities5["学生2"]["性别"]= "女"
	cities5["学生2"]["地址"]= "天通苑"
	cities5["学生3"]["名字"]= "杰克"
	cities5["学生3"]["性别"]= "男"
	cities5["学生3"]["地址"]= "青年汇"
	fmt.Println(cities5["学生3"])
}
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter08\deom02\main.go
// 方式1 map[a:golang]
// 方式2 map[q:你好]
// 方式3 map[node1:晨曦 node2:上海]
// 方式4 map[八戒:九齿钉耙 哪吒:火尖枪 悟空:金箍棒 沙和尚:降魔杖]
// map[名字:杰克 地址:青年汇 性别:男]