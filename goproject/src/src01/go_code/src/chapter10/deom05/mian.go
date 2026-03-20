package main
import (
	"fmt"
)
type A struct{
	Name string
	age int
}
type B struct{
	A
	int  //将内置类型直接作为匿名字段
}
type Goods struct{
	Name string
	Price float64
}
type Brand struct{
	Name string
	Address string
}
type TV  struct{
	Goods
	Brand
	
}
type TV2 struct{
	*Goods
	*Brand
}

// type C struct{
// 	A
// 	B
// }
type F struct{
	a A
}
func main(){
//    var R C
//    R.A.Name="小光"
//    R.B.Name="小度"
//    fmt.Println("A=",R.A.Name)
//    fmt.Println("B=",R.B.Name)
//    var  T F
//    T.a.Name="小曦" //如果结构体嵌套一个有名的结构体，这种模式是组合，那么调用方法或字段时必须带上结构体名字
//    T.a.age=19
//    fmt.Printf("%v的年龄是%v\n",T.a.Name,T.a.age)
    // T:=B{A{"小鹿",18},"小🐺"}
	// fmt.Println(T)
	// tv1:=TV{Goods{"电视剧001",5000},Brand{"海尔","山东"}}
	// tv2:=TV{
	// 	Goods{
	// 		Price: 5000.99,
	// 		Name: "电视机2",
	// 	},
	// 	Brand{
	// 		Name: "小米",
	// 		Address: "北京",
	// 	},
	// }
	// fmt.Println(tv1)
	// fmt.Println(tv2)
	// tv3:=TV2{&Goods{"电视剧003",5000},&Brand{"东芝","山东"}}
	// fmt.Println("tv3",*tv3.Goods,*tv3.Brand)
	y :=B{}
	fmt.Println(y.int)
	y.int=8//修改匿名内置类型字段默认值
	y.age=9
	y.Name="tom"
	fmt.Println(y)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter10\deom05\mian.go
// 0
// {{tom 9} 8}