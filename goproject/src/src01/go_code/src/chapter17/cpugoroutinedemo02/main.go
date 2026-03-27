package main
import(
	"fmt"
	"runtime"
)
func main(){
	//获取当前系统CPU 个数
	num:=runtime.NumCPU()
	//这里设置num-1的CPU运行Go程序
	runtime.GOMAXPROCS(num)
	fmt.Println("cpu个数=",num)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\cpugoroutinedemo02\main.go
// cpu个数= 8