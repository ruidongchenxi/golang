
package main
import (
	"fmt"
)
func main(){
	var breakAgain bool
	//外层
	for x := 0; x < 10; x++{
		//内层
		for y := 0; y < 10;y++{
			if y == 2{
				//设置退出标记
				breakAgain = true
				//退出本次循环
				break
			}
			fmt.Println(y)
		}
		fmt.Println("www",x)
		if breakAgain{
			break
		}
		fmt.Println("chenxi",x)
	} 
}
/*
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\流程控制\使用goto退出多层循环\main0.go
0
1
www 0
构建外层循环
构建内层循环
当y==2 时需要退出所有for循环
默认情况只能一层一层退出，为此需要设置一个状态isbreak，需要退出时，设置这个变量为true
使用break退出循环，执行后代码进入20行输出
进行判断 退出循环
*/