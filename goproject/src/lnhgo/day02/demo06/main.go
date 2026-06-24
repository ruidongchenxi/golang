package main
import(
	"fmt"
)
func main(){
	//map 是key和value（值）的无序，注意查询方便
	var courseMap = map[string]string{
		"go":"go工程师",
		"python":"python工程师",
		"gin":"gin深入学习",
	}
	//打印map某值
	// fmt.Println(courseMap["go"])
	// //放值
	// courseMap["mysql"]="mysql工程师"
	// fmt.Println(courseMap)
	// //声明没有初始化不能直接使用它是nil，要想往里面存值，必须要初始化操作故：var m1 map[string]string 操作失败
	// var m1= map[string]string{} //{}里可以不写内存必须要有，代表初始化
	// m1["s1"]="T"
	// fmt.Println(m1)
	// //使用make 方法初始化;make内置函数主要初始化切片、map、channel
	// var m2= make(map[string]string,3)
	// m2["1"]="a"
	// fmt.Println(m2)
	// //使用make方法初始化2
	// var m3=make(map[string]string)
	// m3["2"]="R"
	// fmt.Println(m3)
	// for i,v:=range courseMap{
	// 	fmt.Println(i,v)
	// }
	// //只遍历key是还可以使用一下方式
	// for v:=range courseMap{
	// 	fmt.Println(v)
	// }
	//判断map里是否存在key
	courseMap["java"]=""
	d,i:=courseMap["java"]//java 的value赋值给d
	if i{
		fmt.Println("map里存在java这个key")
		fmt.Println(d)
	}else{
		fmt.Println("map里不存在java这个key")
	}
	courseMap["java"]="java工程师"
	//d,i=courseMap["java"]//java 的value赋值给d
	if d,i=courseMap["java"]; i{
		fmt.Println("map里存在java这个key")
		fmt.Println(d)
	}else{
		fmt.Println("map里不存在java这个key")
	}
	//删除某个元素,删除不存在的key也不报错；map不是线程安全的
	delete(courseMap,"grpc")
	fmt.Println(courseMap)
	
}