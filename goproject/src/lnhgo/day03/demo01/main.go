package main
import (
	"fmt"
)
/*type关键字
1.定义结构体
2.定义接口
3.定义别名
*/
type MyInt int//类型定义；可以绑定方法
func main(){
	type Rint int//类型定义；通过已有类型自定义类型
	type eint = int//"="号是别名
	var e Rint
	var r MyInt = 5
	var t int = 6
	var w eint =9
	fmt.Printf("%T\n",r)	
	fmt.Printf("%T\r\n",e)
	fmt.Println(r+MyInt(t))
	fmt.Println(w+t)//在编译的时候，类型别名会被直接替换成对应类型（int）
	fmt.Printf("%T\n",w) //查看别名类型
	var a interface{} = "abc"
	switch a.(type){
	case string:
		fmt.Println("string")
	}

}