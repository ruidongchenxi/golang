package main
import "fmt"
func main(){
	//定义map
	var a map[int]string
	//只声明map内存是没有分配
	//必须通过make 函数初始化
	a = make(map[int]string,10)//可以存放10个键值对
	//将键值对存入map中
	a[2009] = "chenxi"
	a[2006] = "chenxi-1"
	a[2345] = "erdt"
	a[2006] = "ed"
	a[2007] = "chenxi"
	//输出集合
	fmt.Println(a)
	//方式2
	b := make(map[int]string)
	b[12] = "cf"
	b[13] = "ds"
	fmt.Println(b)
	//方式3
	c := map[int]string{
		200067 : "golang",
		500989 : "php",
		100000 : "sql" ,
	}
	c[345679] = "kubernetes"
	fmt.Println(c)
	
}