package main
import(
	"fmt"
)
func main() {
	/* iota，特色常量可以被编译器修改的常量
	| 常量名   | iota值 | 表达式    | 结果   |
	| ----- | ----- | ------ | ---- |
	| ERR1  | 0     | iota+1 | 1    |
	| ERR2  | 1     | iota+1 | 2    |
	| ERR25 | 2     | "ha"   | "ha" |
	| ERR4  | 3     | "ha"   | "ha" |
	| ERR42 | 4     | "ha"   | "ha" |
	| ERR5  | 5     | iota   | 5    |
	显示恢复后续自增；自增类型默认int类型
	iota能简化const类型的变量
	每次出现const 的时候iota 就归0

	
	*/
	const (
		ERR1 = iota+1//iota=0 -> 1
		ERR2 //iota=0 -> 1
		ERR25 = "ha"//iota=2
		ERR4 //复制上一行表达式
		ERR42//复制上一行表达式
		ERR45=100//iota5
		ERR5 = iota//显示恢复后续自增

	)
	const (
		ERRNEW1=iota //从0开始
	)

	fmt.Println(ERR1,ERR2,ERR25,ERR4,ERR5)
	fmt.Println(ERRNEW1)
}
